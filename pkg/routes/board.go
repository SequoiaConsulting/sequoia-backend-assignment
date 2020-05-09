package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/model"

	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/service"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type boardController struct {
	svc    service.BoardService
	logger zerolog.Logger
}

// RegisterBoardRoutes registers board related routes on the given router
func RegisterBoardRoutes(router *mux.Router, svc service.BoardService) {
	controller := &boardController{
		svc:    svc,
		logger: log.With().Str("controller", "board").Logger(),
	}

	router.Path("/boards").Methods(http.MethodPost).HandlerFunc(controller.createBoard)
	boardRouter := router.PathPrefix("/boards/{boardID:[0-9]+}").Subrouter()
	boardRouter.Path("/archive").Methods(http.MethodGet).HandlerFunc(controller.archiveBoard(true))
	boardRouter.Path("/unarchive").Methods(http.MethodGet).HandlerFunc(controller.archiveBoard(false))
	boardRouter.Path("/users").Methods(http.MethodPut).HandlerFunc(controller.assignOrRemoveUser(false))
	boardRouter.Path("/users").Methods(http.MethodDelete).HandlerFunc(controller.assignOrRemoveUser(true))
	boardRouter.Path("/users").Queries("limit", "{[0-9]+}").Methods(http.MethodGet).HandlerFunc(controller.getUsers)
	boardRouter.Methods(http.MethodGet).HandlerFunc(controller.getBoard)
}

// swagger:parameters createBoard
type createBoardPrams struct {
	// in: body
	Body struct {
		Name string `json:"name"`
	}
}

// swagger:route POST /boards board createBoard
//
// Creates a new board and sets the current user as its admin
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		201: created
//		400: errored
//		403: errored
//		500: errored
func (controller *boardController) createBoard(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(service.SessionClaims{}).(*service.SessionClaims)
	if !ok {
		controller.logger.Warn().Msg("unable to get auth context")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrInternalServerError)
		return
	}

	if !claims.IsAdmin {
		w.WriteHeader(http.StatusForbidden)
		w.Write(generateErrorResponse("admin level access needed"))
		return
	}

	params := &createBoardPrams{}
	err := parseRequestBody(r.Body, &params.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidBody)
		return
	}

	adminID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	boardID, err := controller.svc.Create(params.Body.Name, uint(adminID))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := &createdResponse{}
	resp.Body.Href = fmt.Sprintf("/boards/%d", boardID)
	w.WriteHeader(http.StatusCreated)
	w.Write(mustJSONMarshal(resp))
}

// swagger:parameters getBoard archiveBoard unarchiveBoard assignUser removeUser listUser
type _ struct {
	// in: path
	BoardID uint `json:"boardID"`
}

// Indicates that user is assigned to the board
// swagger:response getBoard
type _ struct{}

// swagger:route GET /boards/{boardID} board getBoard
//
// Check if a logged in user has access to the given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		200: getBoard
//		400: errored
//		404: errored
//		500: errored
func (controller *boardController) getBoard(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(service.SessionClaims{}).(*service.SessionClaims)
	if !ok {
		controller.logger.Warn().Msg("unable to get auth context")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrInternalServerError)
		return
	}

	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	params := mux.Vars(r)
	boardID, err := strconv.ParseUint(params["boardID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid board_id"))
		return
	}

	err = controller.svc.IsUserAssignedToBoard(uint(boardID), uint(userID))
	if err != nil {
		if err == model.ErrBoardUserNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Successfully archived
// swagger:response archiveBoard
type _ struct{}

// Successfully unarchived
// swagger:response unarchiveBoard
type _ struct{}

// swagger:route GET /boards/{boardID}/archive board archiveBoard
//
// Archive board with given ID
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		202: archiveBoard
//		400: errored
//		404: errored
//		500: errored

// swagger:route GET /boards/{boardID}/unarchive board unarchiveBoard
//
// Unarchive board with given ID
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		202: unarchiveBoard
//		400: errored
//		404: errored
//		500: errored
func (controller *boardController) archiveBoard(archived bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(service.SessionClaims{}).(*service.SessionClaims)
		if !ok {
			controller.logger.Warn().Msg("unable to get auth context")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(ErrInternalServerError)
			return
		}

		if !claims.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			w.Write(generateErrorResponse("admin level access neeeded"))
			return
		}

		userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
		params := mux.Vars(r)
		boardID, err := strconv.ParseUint(params["boardID"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(generateErrorResponse("invalid board_id"))
			return
		}

		err = controller.svc.SetArchived(uint(boardID), uint(userID), archived)
		if err != nil {
			if err == model.ErrBoardNotFound {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write(generateErrorResponse(err.Error()))
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

// swagger:parameters assignUser removeUser
type assignUserParams struct {
	// in: body
	Body struct {
		UserID uint `json:"user_id"`
	}
}

// Indicates operation success
// swagger:response assignOrRemoveUser
type _ struct{}

// swagger:route PUT /boards/{boardID}/users board user assignUser
//
// Adds the provided user to the given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		201: assignOrRemoveUser
//		400: errored
//		404: errored
//		500: errored

// swagger:route DELETE /boards/{boardID}/users board user removeUser
//
// Removes the provided user to the given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		201: assignOrRemoveUser
//		400: errored
//		404: errored
//		500: errored

func (controller *boardController) assignOrRemoveUser(remove bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(service.SessionClaims{}).(*service.SessionClaims)
		if !ok {
			controller.logger.Warn().Msg("unable to get auth context")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(ErrInternalServerError)
			return
		}

		if !claims.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			w.Write(generateErrorResponse("admin level access neeeded"))
			return
		}

		adminUserID, _ := strconv.ParseUint(claims.Subject, 10, 64)
		params := mux.Vars(r)
		boardID, err := strconv.ParseUint(params["boardID"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(generateErrorResponse("invalid board_id"))
			return
		}

		bodyParams := &assignUserParams{}
		err = parseRequestBody(r.Body, &bodyParams.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ErrInvalidBody)
			return
		}

		if remove {
			err = controller.svc.RemoveUser(uint(boardID), bodyParams.Body.UserID, uint(adminUserID))
		} else {
			err = controller.svc.AssignUser(uint(boardID), bodyParams.Body.UserID, uint(adminUserID))
		}

		if err != nil {
			if err == model.ErrInternalServerError {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.Write(generateErrorResponse(err.Error()))
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// A list of user resources with their location
// swagger:response listUser
type listUserResponse struct {
	// in: query
	// Minimum: 1
	// Maximum: 50
	Limit uint `json:"limit"`
	// in: body
	Body struct {
		Hrefs []string `json:"hrefs"`
	}
}

// swagger:route GET /boards/{boardID}/users board user listUser
//
// Lists all users assigned to a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		200: listUser
//		400: errored
//		404: errored
//		500: errored
func (controller *boardController) getUsers(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(service.SessionClaims{}).(*service.SessionClaims)
	if !ok {
		controller.logger.Warn().Msg("unable to get auth context")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrInternalServerError)
		return
	}

	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	params := mux.Vars(r)
	boardID, err := strconv.ParseUint(params["boardID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid board_id"))
		return
	}

	limit, err := strconv.ParseUint(r.FormValue("limit"), 10, 64)
	if err != nil || limit < 1 || limit > 50 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid limit"))
		return
	}

	ids, err := controller.svc.ListUsers(uint(boardID), uint(userID), uint(limit))
	if err != nil {
		if err == model.ErrBoardUserNotFound {
			w.WriteHeader(http.StatusForbidden)
			w.Write(generateErrorResponse("user is not assigned to the board"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(generateErrorResponse(err.Error()))
		}
		return
	}

	resp := &listUserResponse{}
	for _, id := range ids {
		resp.Body.Hrefs = append(resp.Body.Hrefs, fmt.Sprintf("/users/%d", id))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mustJSONMarshal(resp.Body))
}
