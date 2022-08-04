package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// create role
type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	Id         uuid.UUID `json:"id" gorm:"primary_key;not null"`
	Name       string    `json:"name" gorm:"not null" validate:"required"`
	Email      string    `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password   string    `json:"password" gorm:"not null" validate:"required,min=8,max=100"`
	Role       Role      `json:"role" gorm:"not null"`
	Otp        string    `json:"otp" gorm:"not null"`
	OtpExpired int64     `json:"otp_expiry" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
