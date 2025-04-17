package models

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID        string         `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);unique" json:"email"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	Address   string         `gorm:"type:text" json:"address"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // soft delete

	// Relasi
	Orders []Order `gorm:"foreignKey:CustomerID" json:"orders"`
}

type CustomerCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type CustomerUpdateRequest struct {
	ID      string `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}
