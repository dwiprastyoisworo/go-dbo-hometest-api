package services

import (
	"context"
	"errors"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/repositories"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/config"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	repo      repositories.RepositoryImpl[models.User]
	repoUser  repositories.UserRepositoryInterface
	validator *validator.Validate
	db        *gorm.DB
	cfg       config.AppConfig
}

func NewUserService(repo repositories.RepositoryImpl[models.User], repoUser repositories.UserRepositoryInterface, validator *validator.Validate, db *gorm.DB, cfg config.AppConfig) *UserService {
	return &UserService{repo: repo, repoUser: repoUser, validator: validator, db: db, cfg: cfg}
}

type UserServiceInterface interface {
	Login(ctx context.Context, payload models.LoginRequest) (*models.LoginResponse, error)
	Register(ctx context.Context, payload models.RegisterRequest) error
}

func (u UserService) Login(ctx context.Context, payload models.LoginRequest) (*models.LoginResponse, error) {
	err := utils.ValidateStruct(payload, u.validator)
	if err != nil {
		return nil, err
	}

	db := u.db.WithContext(ctx)
	userData, err := u.repoUser.GetUserByUsername(db, payload.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = utils.ComparePassword(userData.PasswordHash, payload.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userData.ID,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(u.cfg.App.JwtExpires))),
	})

	tokenString, err := token.SignedString([]byte(u.cfg.App.JwtSecret))

	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{Token: tokenString}, nil

}

func (u UserService) Register(ctx context.Context, payload models.RegisterRequest) error {
	// validation struct
	err := utils.ValidateStruct(payload, u.validator)
	if err != nil {
		return err
	}

	db := u.db.WithContext(ctx)

	// check username
	err = u.checkUsernameExists(db, payload.Username)
	if err != nil {
		return err
	}

	// generate hash password
	passwordHash, err := utils.GeneratePasswordHash(payload.Password)
	if err != nil {
		return err
	}
	// set payload user data
	user := map[string]interface{}{
		"id":            uuid.NewString(),
		"email":         payload.Email,
		"username":      payload.Username,
		"password_hash": passwordHash,
		"role":          payload.Role,
	}

	// insert data to user table
	err = u.repo.Create(db, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserService) checkUsernameExists(db *gorm.DB, username string) error {
	users, _ := u.repo.DynamicQuery(db, map[string]string{"username": username})
	if len(users) > 0 {
		return errors.New("username already exists")
	}
	return nil
}
