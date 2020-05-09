package service

import (
	"github.com/jinzhu/gorm"
)

// TaskService declares functions implemented by the task service
type TaskService interface {
}

type taskServiceImpl struct {
	db *gorm.DB
}

var _ TaskService = &taskServiceImpl{}

// NewTaskService returns an instance of task service
func NewTaskService(db *gorm.DB) TaskService {
	return &taskServiceImpl{db}
}

// func (svc *taskServiceImpl) Create(task *model.Task) error {
// 	mapping := &model.BoardUser{BoardID: task.BoardID, UserID: task.OwnerID}
// 	err := mapping.Exists(svc.db)
// 	if err != nil {
// 		return err
// 	}

// 	return task.Create(db)
// }

// func (svc *taskServiceImpl) Delete(boardID, taskID, reqUserID uint) error {
// 	mapping := &model.BoardUser{BoardID: task.BoardID, UserID: task.OwnerID}
// 	err := mapping.Exists(svc.db)
// 	if err != nil {
// 		return err
// 	}

// 	return task.Create(db)
// }
