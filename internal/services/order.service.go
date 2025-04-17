package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/repositories"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/config"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"time"
)

type OrderService struct {
	repoOrder     repositories.RepositoryImpl[models.Order]
	repoOrderItem repositories.RepositoryImpl[models.OrderItem]
	repoCustomer  repositories.RepositoryImpl[models.Customer]
	orderRepo     repositories.OrderRepositoryInterface
	validator     *validator.Validate
	db            *gorm.DB
	cfg           config.AppConfig
}

func NewOrderService(repoOrder repositories.RepositoryImpl[models.Order], repoOrderItem repositories.RepositoryImpl[models.OrderItem], repoCustomer repositories.RepositoryImpl[models.Customer], orderRepo repositories.OrderRepositoryInterface, validator *validator.Validate, db *gorm.DB, cfg config.AppConfig) OrderServiceInterface {
	return &OrderService{repoOrder: repoOrder, repoOrderItem: repoOrderItem, repoCustomer: repoCustomer, orderRepo: orderRepo, validator: validator, db: db, cfg: cfg}
}

type OrderServiceInterface interface {
	GetAll(ctx context.Context, page, pageSize int) ([]models.Order, int64, error)
	GetDetail(ctx context.Context, id string) (*models.Order, error)
	Create(ctx context.Context, payload *models.OrderCreateRequest) error
	Update(ctx context.Context, payload *models.OrderUpdateRequest) error
	Delete(ctx context.Context, id string) error
}

func (o *OrderService) GetAll(ctx context.Context, page, pageSize int) ([]models.Order, int64, error) {
	order, total, err := o.repoOrder.GetAll(o.db.WithContext(ctx), page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	for i := range order {
		orderItems, err := o.orderRepo.GetOrderItemsByOrderID(o.db.WithContext(ctx), order[i].ID)
		if err != nil {
			return nil, 0, err
		}
		order[i].OrderItems = orderItems

		customer, _ := o.repoCustomer.GetByID(o.db.WithContext(ctx), order[i].CustomerID)
		if customer != nil {
			order[i].Customer = *customer
		}
	}

	return order, total, nil
}

func (o *OrderService) GetDetail(ctx context.Context, id string) (*models.Order, error) {
	return o.repoOrder.GetByID(o.db.WithContext(ctx), id)
}

func (o *OrderService) Create(ctx context.Context, payload *models.OrderCreateRequest) error {
	// validation struct
	err := o.validator.Struct(payload)
	if err != nil {
		return err
	}

	tx := o.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	orderId := uuid.New().String()

	var orderItems []models.OrderItem
	for _, item := range payload.OrderItems {
		orderItem := models.OrderItem{
			ID:          uuid.NewString(),
			OrderID:     orderId,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
		}
		orderItems = append(orderItems, orderItem)
	}

	orderDate := time.Now().Format("2006-01-02")
	orderDateParse, err := time.Parse("2006-01-02", orderDate)
	if err != nil {
		log.Print(err.Error())
		return errors.New("failed to create order")
	}

	order := models.Order{
		ID:         orderId,
		CustomerID: payload.CustomerID,
		OrderDate:  orderDateParse,
		Status:     "pending",
		OrderItems: orderItems,
	}

	// Recalculate total amount
	order.RecalculateOrderAmount()

	if err := o.repoOrder.Create(tx, order.OrderToMap()); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create order: %w", err)
	}

	if err := o.repoOrderItem.CreateBatch(tx, order.OrderItemsToMap()); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create order items: %w", err)
	}

	return tx.Commit().Error
}

func (o *OrderService) Update(ctx context.Context, payload *models.OrderUpdateRequest) error {
	// validation struct
	err := o.validator.Struct(payload)
	if err != nil {
		return err
	}

	tx := o.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// get data order from database
	order, err := o.repoOrder.GetByID(tx, payload.ID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get order: %w", err)
	}

	// mapping data orderItem from payload
	var orderItems []models.OrderItem
	for _, item := range payload.OrderItems {
		orderItem := models.OrderItem{
			ID:          item.ID,
			OrderID:     order.ID,
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
		}
		orderItems = append(orderItems, orderItem)
	}

	order.OrderItems = orderItems
	order.RecalculateOrderAmount()

	// update data order
	if err := o.repoOrder.Update(tx, order.OrderToMap()); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update order: %w", err)
	}

	// update data orderItem
	if err := o.repoOrderItem.UpdateBatch(tx, order.OrderItemsToMap()); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update order items: %w", err)
	}

	return tx.Commit().Error
}

func (o *OrderService) Delete(ctx context.Context, id string) error {
	// get data order from database
	order, err := o.repoOrder.GetByID(o.db.WithContext(ctx), id)
	if err != nil {
		return fmt.Errorf("failed to get order: %w", err)
	}

	tx := o.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := o.repoOrder.Delete(tx, order.ID); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return tx.Commit().Error
}
