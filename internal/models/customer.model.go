package models

import "time"

type Customer struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);unique" json:"email"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Address   string    `gorm:"type:text" json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relasi
	Orders []Order `gorm:"foreignKey:CustomerID" json:"orders"`
}
