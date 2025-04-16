package repositories

import (
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

type UserRepositoryInterface interface {
	GetUserByUsername(db *gorm.DB, username string) (*models.User, error)
}

func (u UserRepository) GetUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	var userData models.User
	err := db.Where("username = ?", username).First(&userData).Error
	return &userData, err
}
