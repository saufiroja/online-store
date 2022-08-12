package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id   string `json:"id" gorm:"primary_key;not null"`
	Name string `json:"name" gorm:"not null" validate:"required"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
