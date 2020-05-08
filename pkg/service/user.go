package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/ashutoshgngwr/sequoia-backend-assignment/pkg/model"
)

// SessionClaims declares the claim structure stored in the JWT playload for the session
type SessionClaims struct {
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}

// UserService declares all the functions performed by the user
// service.
type UserService interface {
	// Get just gets the the user
	Get(id uint) (*model.User, error)
	// Create creates a new user in the database and returns its ID.
	Create(*model.User) (uint, error)
	// Login tries to authenticate a user using username and password. Returns a session
	// token on successful login, returns error otherwise.
	Login(email, password string) (string, error)
	// Authenticate authenticates a user using a session token. If the token is valid,
	// it returns a nil error. Returns a non-nil error otherwise.
	Authenticate(token string) (*SessionClaims, error)
	// RevokeSession revokes the provided session.
	RevokeSession(token string) error
}

type userServiceImpl struct {
	db         *gorm.DB
	signingKey []byte
	logger     zerolog.Logger
}

// NewUserService returns a new instance of user service
func NewUserService(db *gorm.DB, jwtSigningKey string) UserService {
	return &userServiceImpl{
		db,
		[]byte(jwtSigningKey),
		log.With().Str("service", "user").Logger(),
	}
}

var _ UserService = &userServiceImpl{}

func (svc *userServiceImpl) Get(id uint) (*model.User, error) {
	user := &model.User{}
	err := user.FindByID(svc.db, id)
	return user, err
}

func (svc *userServiceImpl) Create(user *model.User) (uint, error) {
	if err := user.Validate(); err != nil {
		return 0, err
	}

	user.ID = 0 // ensure ID is 0 when creating a new user
	err := user.Create(svc.db)
	if err == model.ErrDuplicateEntry {
		return 0, errors.New("user already exists")
	}

	return user.ID, err
}

func (svc *userServiceImpl) Login(email, password string) (string, error) {
	user := &model.User{}
	err := user.FindByEmail(svc.db, email)
	if err != nil {
		if err == model.ErrUserNotFound {
			return "", errors.New("invalid login credentials")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid login credentials")
	}

	claims := SessionClaims{
		IsAdmin: *user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			Subject:   string(user.ID),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(svc.signingKey)
	if err != nil {
		svc.logger.Warn().Err(err).Msg("unable to sign jwt token")
		return "", model.ErrInternalServerError
	}

	return signed, nil
}

func (svc *userServiceImpl) keyFunc(token *jwt.Token) (interface{}, error) {
	return svc.signingKey, nil
}

func (svc *userServiceImpl) Authenticate(tokenString string) (*SessionClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SessionClaims{}, svc.keyFunc)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid session token")
	}

	claims, ok := token.Claims.(*SessionClaims)
	if !ok {
		return nil, errors.New("invalid session token")
	}
	return claims, nil
}

func (svc *userServiceImpl) RevokeSession(token string) error {
	claims, err := svc.Authenticate(token)
	if err != nil {
		return errors.New("invalid session token")
	}

	session := &model.RevokedSession{
		Token:   token,
		Expires: time.Unix(claims.ExpiresAt, 0),
	}

	err = session.Create(svc.db)
	if err == model.ErrDuplicateEntry {
		return errors.New("session is already revoked")
	}
	return err
}
