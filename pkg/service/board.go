package service

import (
	"errors"

	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// BoardService declares functions implemented by Board service
type BoardService interface {
	// Creates a board with given name. Returns id on success, error otherwise
	Create(name string, adminID uint) (uint, error)
	// IsUserAssignedToBoard checks if a given user can access the given board.
	// If user has access, it returns a nil error. It returns non-nil error otherwise
	IsUserAssignedToBoard(boardID, userID uint) error
	// SetArchived archives or unarchives the given board if requesting user has permission.
	SetArchived(boardID, reqUserID uint, archived bool) error
	// AssignUser assigns the given user to the board if requesting user has permission.
	AssignUser(boardID, userID, reqUserID uint) error
	// RemoveUser removes the user from the given board if requesting user has permission.
	RemoveUser(boardID, userID, reqUserID uint) error
	// ListUsers lists all users for a board if requesting user is assigned to it
	ListUsers(boardID, reqUserID, limit uint) ([]uint, error)
	// CreateStaus creates a new task status with given details
	CreateStatus(title string, boardID, reqUserID uint) (uint, error)
	// DeleteStatus deletes the task status with given id
	DeleteStatus(statusID, boardID, reqUserID uint) error
	// ListStatus lists all status for a given board
	ListStatus(boardID, reqUserID uint) ([]uint, error)
	// GetStatus gets details of specific status
	GetStatus(statusID, boardID, reqUserID uint) (*model.TaskStatus, error)
}

type boardServiceImpl struct {
	db     *gorm.DB
	logger zerolog.Logger
}

var _ BoardService = &boardServiceImpl{}

// NewBoardService returns a concrete implementation of board service.
func NewBoardService(db *gorm.DB) BoardService {
	return &boardServiceImpl{
		db:     db,
		logger: log.With().Str("service", "board").Logger(),
	}
}

func (svc *boardServiceImpl) Create(name string, adminID uint) (uint, error) {
	board := &model.Board{Name: name, AdminUserID: adminID}
	err := svc.db.Transaction(func(txn *gorm.DB) error {
		err := board.Create(txn)
		if err != nil {
			if err == model.ErrDuplicateEntry {
				return errors.New("board with name already exists")
			}
			return err
		}

		boardUser := &model.BoardUser{BoardID: board.ID, UserID: adminID}
		return boardUser.Create(txn)
	})

	return board.ID, err
}

func (svc *boardServiceImpl) SetArchived(boardID, reqUserID uint, archived bool) error {
	board := &model.Board{}
	err := board.FindByID(svc.db, boardID)
	if err != nil {
		return err
	}

	if board.AdminUserID == reqUserID {
		board.IsArchived = &archived
		return board.Update(svc.db)
	}

	return model.ErrBoardNotFound // mask the forbidden thing?
}

func (svc *boardServiceImpl) IsUserAssignedToBoard(boardID, userID uint) error {
	mapping := &model.BoardUser{BoardID: boardID, UserID: userID}
	return mapping.Exists(svc.db)
}

func (svc *boardServiceImpl) AssignUser(boardID, userID, reqUserID uint) error {
	user := &model.User{}
	err := user.FindByID(svc.db, userID)
	if err != nil {
		return err
	}

	board := &model.Board{}
	err = board.FindByID(svc.db, boardID)
	if err != nil {
		return err
	}

	if board.AdminUserID == reqUserID {
		if *board.IsArchived {
			return errors.New("board is archived")
		}

		mapping := &model.BoardUser{BoardID: boardID, UserID: userID}
		return mapping.Create(svc.db)
	}

	return model.ErrBoardNotFound
}

func (svc *boardServiceImpl) RemoveUser(boardID, userID, reqUserID uint) error {
	board := &model.Board{}
	err := board.FindByID(svc.db, boardID)
	if err != nil {
		return err
	}

	if board.AdminUserID == reqUserID {
		if *board.IsArchived {
			return errors.New("board is archived")
		}

		mapping := &model.BoardUser{BoardID: boardID, UserID: userID}
		return mapping.Delete(svc.db)
	}

	return model.ErrBoardNotFound
}

func (svc *boardServiceImpl) ListUsers(boardID, reqUserID, limit uint) ([]uint, error) {
	boardUser := &model.BoardUser{BoardID: boardID, UserID: reqUserID}
	err := boardUser.Exists(svc.db)
	if err != nil {
		return nil, err
	}

	userIDs := []uint{}
	err = svc.db.Model(&model.BoardUser{}).
		Where(&model.BoardUser{BoardID: boardUser.BoardID}).
		Limit(limit).
		Pluck("user_id", &userIDs).Error

	if err != nil {
		svc.logger.Warn().Err(err).Msg("unable to list board users")
		return nil, model.ErrInternalServerError
	}

	return userIDs, nil
}

func (svc *boardServiceImpl) CreateStatus(title string, boardID, reqUserID uint) (uint, error) {
	boardUser := &model.BoardUser{BoardID: boardID, UserID: reqUserID}
	err := boardUser.Exists(svc.db)
	if err != nil {
		return 0, err
	}

	status := &model.TaskStatus{BoardID: boardID, Title: title}
	err = status.Create(svc.db)
	return status.ID, err
}

func (svc *boardServiceImpl) DeleteStatus(statusID, boardID, reqUserID uint) error {
	return svc.db.Transaction(func(txn *gorm.DB) error {
		boardUser := &model.BoardUser{BoardID: boardID, UserID: reqUserID}
		err := boardUser.Exists(txn)
		if err != nil {
			return err
		}

		status := &model.TaskStatus{Model: gorm.Model{ID: statusID}}
		err = status.Find(txn)
		if err != nil {
			return err
		}

		// set task status to null
		err = txn.Model(&model.Task{}).Where("status_id = ?", status.ID).UpdateColumns(map[string]interface{}{
			"status_id": nil,
		}).Error

		if err != nil {
			svc.logger.Warn().Err(err).Msg("unable to set status_id to null")
			return model.ErrInternalServerError
		}

		return status.Delete(txn)
	})
}

func (svc *boardServiceImpl) ListStatus(boardID, reqUserID uint) ([]uint, error) {
	boardUser := &model.BoardUser{BoardID: boardID, UserID: reqUserID}
	err := boardUser.Exists(svc.db)
	if err != nil {
		return nil, err
	}

	ids := []uint{}
	where := &model.TaskStatus{BoardID: boardID}
	err = svc.db.Model(where).Where(where).Pluck("id", &ids).Error
	if err != nil {
		svc.logger.Warn().Err(err).Msg("unable to exec list status query")
		return nil, model.ErrInternalServerError
	}
	return ids, nil
}

func (svc *boardServiceImpl) GetStatus(statusID, boardID, reqUserID uint) (*model.TaskStatus, error) {
	boardUser := &model.BoardUser{BoardID: boardID, UserID: reqUserID}
	err := boardUser.Exists(svc.db)
	if err != nil {
		return nil, err
	}

	status := &model.TaskStatus{Model: gorm.Model{ID: statusID}}
	err = status.Find(svc.db)
	return status, err
}
