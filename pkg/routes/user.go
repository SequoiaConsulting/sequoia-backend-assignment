package routes

import (
	"fmt"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/model"
	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/service"
)

type userController struct {
	svc    service.UserService
	logger zerolog.Logger
}

// RegisterUserRoutes registers user related routes on the given http.Router
func RegisterUserRoutes(router *mux.Router, svc service.UserService) {
	controller := &userController{svc, log.With().Str("controller", "user").Logger()}
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/signup", controller.signup).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", controller.login).Methods(http.MethodPost)
	userRouter.HandleFunc("/logout", controller.logout).Methods(http.MethodGet)
	userRouter.HandleFunc("/{userID:[0-9]+}", controller.get).Methods(http.MethodGet)
	controller.logger.Debug().Msg("registered routes")
}

// swagger:parameters signupUser
type signupParams struct {
	// in: body
	// required: true
	User struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

// swagger:route POST /users/signup user signupUser
//
// Register a new user with the system.
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Responses:
//		201: created
//		400: errored
//		500: errored
func (controller *userController) signup(w http.ResponseWriter, r *http.Request) {
	params := signupParams{}
	err := parseRequestBody(r.Body, &params.User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidBody)
		return
	}

	err = validation.Validate(&params.User.Password, validation.Required, validation.Length(8, 0))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse(fmt.Sprintf("Password: %s", err.Error())))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(params.User.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrInternalServerError)
		return
	}

	id, err := controller.svc.Create(&model.User{
		Name:     params.User.Name,
		Email:    params.User.Email,
		Password: string(hash),
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := &createdResponse{}
	resp.Body.Href = fmt.Sprintf("/user/%d", id)
	w.WriteHeader(http.StatusAccepted)
	w.Write(mustJSONMarshal(resp))
}

// swagger:parameters loginUser
type loginParams struct {
	// in: body
	// required: true
	Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

// loginReponse returns a session token in Authorization header on successful login
// swagger:response loginResponse
type loginResponse struct {
	// Bearer <token>
	Authorization string
}

// swagger:route POST /users/login user loginUser
//
// Login a user using credentials to issue a session token
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Responses:
//		200: loginResponse
//		400: errored
//		403: errored
//		500: errored
func (controller *userController) login(w http.ResponseWriter, r *http.Request) {
	params := loginParams{}
	err := parseRequestBody(r.Body, &params.Credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidBody)
		return
	}

	token, err := controller.svc.Login(params.Credentials.Email, params.Credentials.Password)
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
}

// indicates successful logout
// swagger:response logoutResponse
type _ struct{}

// swagger:route GET /users/logout user logoutUser
//
// Revokes the current session token for the user
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		201: logoutResponse
//		403: errored
//		500: errored
func (controller *userController) logout(w http.ResponseWriter, r *http.Request) {
	err := controller.svc.RevokeSession(r.Header.Get("Authorization"))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}
		w.Write(generateErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

// swagger:parameters getUser
type getUserParams struct {
	// in: path
	UserID string `json:"userID"`
}

// Returns requested user
// swagger:response getUser
type getUserResponse struct {
	// in: body
	// required: true
	User struct {
		ID      uint   `json:"id"`
		Name    string `json:"name"`
		Email   string `json:"email"`
		IsAdmin bool   `json:"is_admin"`
	}
}

// swagger:route GET /users/{userID} user getUser
//
// Get a user by their ID
//
//	Consumes:
//		- application/json
//	Produces:
//		- application/json
//	Schemes: http, https
//	Security:
//		- api_key
//	Responses:
//		200: getUser
//		404: errored
//		500: errored
func (controller *userController) get(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(service.SessionClaims{}).(*service.SessionClaims)
	if !ok {
		controller.logger.Warn().Msg("unable to get auth context")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrInternalServerError)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["userID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generateErrorResponse("invalid user id"))
		return
	}

	requestingUserID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	user, err := controller.svc.Get(uint(userID), uint(requestingUserID))
	if err != nil {
		if err == model.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Write(generateErrorResponse(err.Error()))
		return
	}

	resp := &getUserResponse{}
	resp.User.ID = user.ID
	resp.User.Name = user.Name
	resp.User.Email = user.Email
	resp.User.IsAdmin = *user.IsAdmin
	w.WriteHeader(http.StatusOK)
	w.Write(mustJSONMarshal(&resp.User))
}
