package repositories

import (
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
}

func NewOrderRepository() OrderRepositoryInterface {
	return &OrderRepository{}
}

type OrderRepositoryInterface interface {
	GetOrderItemsByOrderID(db *gorm.DB, orderId string) ([]models.OrderItem, error)
}

func (o *OrderRepository) GetOrderItemsByOrderID(db *gorm.DB, orderId string) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := db.Where("order_id = ?", orderId).Find(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}
