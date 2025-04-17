package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID          string         `gorm:"type:uuid;primaryKey" json:"id"`
	CustomerID  string         `gorm:"type:uuid;not null" json:"customer_id"`
	OrderDate   time.Time      `gorm:"autoCreateTime" json:"order_date"`
	Status      string         `gorm:"type:varchar(50)" json:"status"` // e.g. pending, paid
	TotalAmount float64        `gorm:"type:decimal(10,2)" json:"total_amount"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // soft delete

	// Relasi
	Customer   Customer    `gorm:"foreignKey:CustomerID;references:ID" json:"customer"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}

type OrderItem struct {
	ID          string  `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID     string  `gorm:"type:uuid;not null" json:"order_id"`
	ProductName string  `gorm:"type:varchar(255)" json:"product_name"`
	Quantity    int     `gorm:"not null" json:"quantity"`
	Price       float64 `gorm:"type:decimal(10,2)" json:"price"`
	Subtotal    float64 `gorm:"type:decimal(10,2)" json:"subtotal"`
}

type OrderCreateRequest struct {
	CustomerID string                   `json:"customer_id" validate:"required"`
	Status     string                   `json:"status" validate:"required"`
	OrderItems []OrderItemCreateRequest `json:"order_items" validate:"required,dive"`
}

type OrderItemCreateRequest struct {
	ProductName string  `json:"product_name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type OrderUpdateRequest struct {
	ID         string                   `json:"id"`
	CustomerID string                   `json:"customer_id" validate:"required"`
	OrderItems []OrderItemUpdateRequest `json:"order_items" validate:"required,dive"`
}

type OrderItemUpdateRequest struct {
	ID          string  `json:"id" validate:"required"`
	ProductName string  `json:"product_name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

func (o *Order) RecalculateOrderAmount() {
	// Recalculate subtotal for each order item
	var totalAmount float64
	for i := range o.OrderItems {
		o.OrderItems[i].Subtotal = float64(o.OrderItems[i].Quantity) * o.OrderItems[i].Price
		totalAmount += o.OrderItems[i].Subtotal
	}
	// Update total amount in the order
	o.TotalAmount = totalAmount
}

func (o *Order) OrderToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           o.ID,
		"customer_id":  o.CustomerID,
		"order_date":   o.OrderDate,
		"status":       o.Status,
		"total_amount": o.TotalAmount,
	}
}

func (o *Order) OrderItemsToMap() []map[string]interface{} {
	orderItemMaps := make([]map[string]interface{}, len(o.OrderItems))
	for i, item := range o.OrderItems {
		orderItemMaps[i] = map[string]interface{}{
			"id":           item.ID,
			"order_id":     item.OrderID,
			"product_name": item.ProductName,
			"quantity":     item.Quantity,
			"price":        item.Price,
			"subtotal":     item.Subtotal,
		}
	}
	return orderItemMaps
}
