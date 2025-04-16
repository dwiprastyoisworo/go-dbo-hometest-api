package models

import "time"

type User struct {
	ID           string    `gorm:"type:uuid;primaryKey" json:"id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	Role         string    `gorm:"type:varchar(255)" json:"role"` // e.g. admin, user
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type LoginLog struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    string    `gorm:"type:uuid;not null" json:"user_id"`
	IPAddress string    `gorm:"type:varchar(255)" json:"ip_address"`
	UserAgent string    `gorm:"type:text" json:"user_agent"`
	LoginTime time.Time `gorm:"autoCreateTime" json:"login_time"`

	// Relasi ke User
	User User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required"`
}
