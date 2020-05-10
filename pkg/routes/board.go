package routes

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

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
	router.Path("/boards/{boardID:[0-9]+}").Methods(http.MethodGet).HandlerFunc(controller.getBoard)
	boardRouter := router.PathPrefix("/boards/{boardID:[0-9]+}").Subrouter()
	boardRouter.Path("/archive").Methods(http.MethodGet).HandlerFunc(controller.archiveBoard(true))
	boardRouter.Path("/unarchive").Methods(http.MethodGet).HandlerFunc(controller.archiveBoard(false))
	boardRouter.Path("/users").Methods(http.MethodPut).HandlerFunc(controller.assignOrRemoveUser(false))
	boardRouter.Path("/users").Methods(http.MethodDelete).HandlerFunc(controller.assignOrRemoveUser(true))
	boardRouter.Path("/users").Queries("limit", "{[0-9]+}").Methods(http.MethodGet).HandlerFunc(controller.getUsers)
	boardRouter.Path("/users").Methods(http.MethodGet).HandlerFunc(controller.getUsers)
	boardRouter.Path("/status").Methods(http.MethodPost).HandlerFunc(controller.createStatus)
	boardRouter.Path("/status").Queries("limit", "{[0-9]+}").Methods(http.MethodGet).HandlerFunc(controller.listStatus)
	boardRouter.Path("/status").Methods(http.MethodGet).HandlerFunc(controller.listStatus)
	boardRouter.Path("/status/{statusID:[0-9]+}").Methods(http.MethodDelete).HandlerFunc(controller.deleteStatus)
	boardRouter.Path("/status/{statusID:[0-9]+}").Methods(http.MethodGet).HandlerFunc(controller.getStatus)
	boardRouter.Path("/tasks").Methods(http.MethodPost).HandlerFunc(controller.createTask)
	boardRouter.Path("/tasks").Queries("status_id", "{[0-9]+}", "limit", "{[0-9]+}").Methods(http.MethodGet).HandlerFunc(controller.listTask)
	boardRouter.Path("/tasks").Methods(http.MethodGet).HandlerFunc(controller.listTask)
	boardRouter.Path("/tasks/{taskID:[0-9]+}").Methods(http.MethodDelete).HandlerFunc(controller.deleteTask)
	boardRouter.Path("/tasks/{taskID:[0-9]+}").Methods(http.MethodPut).HandlerFunc(controller.updateTask)
	boardRouter.Path("/tasks/{taskID:[0-9]+}").Methods(http.MethodGet).HandlerFunc(controller.getTask)
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

// swagger:parameters getBoard archiveBoard unarchiveBoard assignUser removeUser listUser createStatus deleteStatus getStatus listStatus createTask updateTask deleteTask listTask getTask
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

// swagger:parameters listUser listStatus listTask
type _ struct {
	// in: query
	Limit uint `json:"limit"`
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
//		200: listResponse
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
	if err != nil {
		limit = 10 // default
	}

	if limit < 1 || limit > 50 {
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

	resp := &listResponse{}
	for _, id := range ids {
		resp.Body.Hrefs = append(resp.Body.Hrefs, fmt.Sprintf("/users/%d", id))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mustJSONMarshal(resp.Body))
}

// swagger:parameters createStatus
type createStatusParams struct {
	// in: body
	Body struct {
		Title string `json:"title"`
	}
}

// swagger:route POST /boards/{boardID}/status board status createStatus
//
// Create a new status for a given board
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
//		500: errored
func (controller *boardController) createStatus(w http.ResponseWriter, r *http.Request) {
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

	bodyParams := &createStatusParams{}
	err = parseRequestBody(r.Body, &bodyParams.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidBody)
		return
	}

	statusID, err := controller.svc.CreateStatus(bodyParams.Body.Title, uint(boardID), uint(userID))
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
	resp.Body.Href = fmt.Sprintf("/boards/%d/status/%d", boardID, statusID)
	w.WriteHeader(http.StatusCreated)
	w.Write(mustJSONMarshal(resp.Body))
}

// swagger:route GET /boards/{boardID}/status board status listStatus
//
// List all status for a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		200: listResponse
//		400: errored
//		500: errored
func (controller *boardController) listStatus(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		limit = 10 // default
	}

	if limit < 1 || limit > 50 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid limit"))
		return
	}

	statusIDs, err := controller.svc.ListStatus(uint(boardID), uint(userID))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := &listResponse{}
	for _, id := range statusIDs {
		resp.Body.Hrefs = append(resp.Body.Hrefs, fmt.Sprintf("/boards/%d/status/%d", boardID, id))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mustJSONMarshal(resp.Body))
}

// swagger:parameters deleteStatus getStatus
type _ struct {
	// in: path
	StatusID string `json:"statusID"`
}

// status deleted successfully
// swagger:response deleteStatus
type _ struct{}

// swagger:route DELETE /boards/{boardID}/status/{statusID} board status deleteStatus
//
// Delete status for a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		202: deleteStatus
//		400: errored
//		500: errored
func (controller *boardController) deleteStatus(w http.ResponseWriter, r *http.Request) {
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

	statusID, err := strconv.ParseUint(params["statusID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid status_id"))
		return
	}

	err = controller.svc.DeleteStatus(uint(statusID), uint(boardID), uint(userID))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// task status details
// swagger:response getStatus
type getStatusResponse struct {
	// in: body
	Body struct {
		ID      uint   `json:"id"`
		Title   string `json:"title"`
		BoardID uint   `json:"board_id"`
	}
}

// swagger:route GET /boards/{boardID}/status/{statusID} board status getStatus
//
// Get status details for a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		200: getStatus
//		400: errored
//		500: errored
func (controller *boardController) getStatus(w http.ResponseWriter, r *http.Request) {
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

	statusID, err := strconv.ParseUint(params["statusID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid status_id"))
		return
	}

	status, err := controller.svc.GetStatus(uint(statusID), uint(boardID), uint(userID))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := &getStatusResponse{}
	resp.Body.ID = status.ID
	resp.Body.Title = status.Title
	resp.Body.BoardID = status.BoardID
	w.WriteHeader(http.StatusAccepted)
	w.Write(mustJSONMarshal(resp.Body))
}

// status_id, assignee_id and due_date are optional
// swagger:parameters createTask updateTask
type createTaskParams struct {
	// in: body
	Body struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		StatusID    uint      `json:"status_id,omitempty"`
		DueDate     time.Time `json:"due_date,omitempty"`
		AssigneeID  uint      `json:"assignee_id,omitempty"`
	}
}

// swagger:route POST /boards/{boardID}/tasks board task createTask
//
// Create a new task on a given board
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
//		500: errored
func (controller *boardController) createTask(w http.ResponseWriter, r *http.Request) {
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

	bodyParams := &createTaskParams{}
	err = parseRequestBody(r.Body, &bodyParams.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidBody)
		return
	}

	task := &model.Task{
		Title:       bodyParams.Body.Title,
		Description: bodyParams.Body.Description,
		BoardID:     uint(boardID),
		OwnerID:     uint(userID),
	}

	if bodyParams.Body.AssigneeID > 0 {
		assigner := uint(userID)
		task.AssigneeID = &bodyParams.Body.AssigneeID
		task.AssignerID = &assigner
	}

	if bodyParams.Body.StatusID > 0 {
		task.StatusID = &bodyParams.Body.StatusID
	}

	if !reflect.DeepEqual(bodyParams.Body.DueDate, time.Time{}) {
		task.DueDate = &bodyParams.Body.DueDate
	}

	taskID, err := controller.svc.CreateTask(task)
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := createdResponse{}
	resp.Body.Href = fmt.Sprintf("/boards/%d/tasks/%d", boardID, taskID)
	w.WriteHeader(http.StatusCreated)
	w.Write(mustJSONMarshal(resp.Body))
}

// swagger:parameters listTask
type _ struct {
	// list tasks for this status only
	// in: query
	StatusID uint `json:"status_id"`
}

// swagger:route GET /boards/{boardID}/tasks board task listTask
//
// List tasks on a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		200: listResponse
//		400: errored
//		500: errored
func (controller *boardController) listTask(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		limit = 10 // default
	}

	if limit < 1 || limit > 50 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid limit"))
		return
	}

	statusID, err := strconv.ParseUint(r.FormValue("status_id"), 10, 64)
	if r.FormValue("status_id") != "" && (err != nil || statusID < 1) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid status id"))
		return
	}

	taskIDs, err := controller.svc.ListTask(uint(boardID), uint(userID), uint(statusID), uint(limit))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := listResponse{}
	for _, id := range taskIDs {
		resp.Body.Hrefs = append(resp.Body.Hrefs, fmt.Sprintf("/boards/%d/tasks/%d", boardID, id))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mustJSONMarshal(resp.Body))
}

// swagger:parameters updateTask deleteTask getTask
type _ struct {
	// in: path
	TaskID uint `json:"taskID"`
}

// Successful update
// swagger:response updateTaskResponse
type _ struct{}

// swagger:route PUT /boards/{boardID}/tasks/{taskID} board task updateTask
//
// Update task on a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		202: updateTaskResponse
//		400: errored
//		500: errored
func (controller *boardController) updateTask(w http.ResponseWriter, r *http.Request) {
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

	taskID, err := strconv.ParseUint(params["taskID"], 10, 64)
	if err != nil || taskID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid taskID"))
		return
	}

	bodyParams := &createTaskParams{}
	err = parseRequestBody(r.Body, &bodyParams.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidBody)
		return
	}

	task := &model.Task{
		Title:       bodyParams.Body.Title,
		Description: bodyParams.Body.Description,
		BoardID:     uint(boardID),
	}

	task.ID = uint(taskID)
	if bodyParams.Body.AssigneeID > 0 {
		assigner := uint(userID)
		task.AssigneeID = &bodyParams.Body.AssigneeID
		task.AssignerID = &assigner
	}

	if bodyParams.Body.StatusID > 0 {
		task.StatusID = &bodyParams.Body.StatusID
	}

	if !reflect.DeepEqual(bodyParams.Body.DueDate, time.Time{}) {
		task.DueDate = &bodyParams.Body.DueDate
	}

	err = controller.svc.UpdateTask(task)
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// Successful delete
// swagger:response deleteTaskResponse
type _ struct{}

// swagger:route DELETE /boards/{boardID}/tasks/{taskID} board task deleteTask
//
// Delete task on a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		202: deleteTaskResponse
//		400: errored
//		500: errored
func (controller *boardController) deleteTask(w http.ResponseWriter, r *http.Request) {
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

	taskID, err := strconv.ParseUint(params["taskID"], 10, 64)
	if err != nil || taskID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid taskID"))
		return
	}

	err = controller.svc.DeleteTask(uint(taskID), uint(boardID), uint(userID))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// Returned task
// swagger:response getTask
type getTaskResponse struct {
	// in: body
	Body struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		OwnerID     uint       `json:"owner_id"`
		DueDate     *time.Time `json:"due_date,omitempty"`
		AssigneeID  uint       `json:"assignee_id,omitempty"`
		AssignerID  uint       `json:"assigner_id,omitempty"`
		StatusID    uint       `json:"status_id,omitempty"`
	}
}

// swagger:route GET /boards/{boardID}/tasks/{taskID} board task getTask
//
// Get task on a given board
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		200: getTask
//		400: errored
//		404: errored
//		500: errored
func (controller *boardController) getTask(w http.ResponseWriter, r *http.Request) {
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

	taskID, err := strconv.ParseUint(params["taskID"], 10, 64)
	if err != nil || taskID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid taskID"))
		return
	}

	task, err := controller.svc.GetTask(uint(taskID), uint(boardID), uint(userID))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else if err == model.ErrTaskNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := &getTaskResponse{}
	resp.Body.Title = task.Title
	resp.Body.Description = task.Description
	resp.Body.OwnerID = task.OwnerID
	if task.AssigneeID != nil {
		resp.Body.AssigneeID = *task.AssigneeID
	}

	if task.AssignerID != nil {
		resp.Body.AssignerID = *task.AssignerID
	}

	if task.DueDate != nil {
		resp.Body.DueDate = task.DueDate
	}

	if task.StatusID != nil {
		resp.Body.StatusID = *task.StatusID
	}

	w.WriteHeader(http.StatusOK)
	w.Write(mustJSONMarshal(resp.Body))
}
