package models

import "time"

type Order struct {
	ID          string    `gorm:"type:uuid;primaryKey" json:"id"`
	CustomerID  string    `gorm:"type:uuid;not null" json:"customer_id"`
	OrderDate   time.Time `gorm:"autoCreateTime" json:"order_date"`
	Status      string    `gorm:"type:varchar(50)" json:"status"` // e.g. pending, paid
	TotalAmount float64   `gorm:"type:decimal(10,2)" json:"total_amount"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

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

	// Relasi
	Order Order `gorm:"foreignKey:OrderID;references:ID" json:"order"`
}
