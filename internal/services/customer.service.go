package services

import (
	"context"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/repositories"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/config"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerService struct {
	repo      repositories.RepositoryImpl[models.Customer]
	validator *validator.Validate
	db        *gorm.DB
	cfg       config.AppConfig
}

func NewCustomerService(repo repositories.RepositoryImpl[models.Customer], validator *validator.Validate, db *gorm.DB, cfg config.AppConfig) CustomerServiceInterface {
	return &CustomerService{repo: repo, validator: validator, db: db, cfg: cfg}
}

type CustomerServiceInterface interface {
	GetAll(ctx context.Context, page, pageSize int) ([]models.Customer, int64, error)
	GetDetail(ctx context.Context, id string) (*models.Customer, error)
	Create(ctx context.Context, customer *models.CustomerCreateRequest) error
	Update(ctx context.Context, customer *models.CustomerUpdateRequest) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, search string, page, pageSize int) ([]models.Customer, int64, error)
}

func (c *CustomerService) GetAll(ctx context.Context, page, pageSize int) ([]models.Customer, int64, error) {
	return c.repo.GetAll(c.db.WithContext(ctx), page, pageSize)
}

func (c *CustomerService) GetDetail(ctx context.Context, id string) (*models.Customer, error) {
	return c.repo.GetByID(c.db.WithContext(ctx), id)
}

func (c *CustomerService) Create(ctx context.Context, customer *models.CustomerCreateRequest) error {
	// validation struct
	err := utils.ValidateStruct(customer, c.validator)
	if err != nil {
		return err
	}

	payload := map[string]interface{}{
		"id":      uuid.NewString(),
		"name":    customer.Name,
		"email":   customer.Email,
		"phone":   customer.Phone,
		"address": customer.Address,
	}

	return c.repo.Create(c.db.WithContext(ctx), payload)
}

func (c *CustomerService) Update(ctx context.Context, customer *models.CustomerUpdateRequest) error {
	// validation struct
	err := utils.ValidateStruct(customer, c.validator)
	if err != nil {
		return err
	}

	payload := map[string]interface{}{
		"id":      customer.ID,
		"name":    customer.Name,
		"email":   customer.Email,
		"phone":   customer.Phone,
		"address": customer.Address,
	}

	return c.repo.Update(c.db.WithContext(ctx), payload)

}

func (c *CustomerService) Delete(ctx context.Context, id string) error {
	return c.repo.Delete(c.db.WithContext(ctx), id)
}

func (c *CustomerService) Search(ctx context.Context, search string, page, pageSize int) ([]models.Customer, int64, error) {
	db := c.db.WithContext(ctx)
	payload := map[string]string{
		"name": search,
	}
	return c.repo.Search(db, payload, page, pageSize)
}
