package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text;not null"`
	Stock       int    `json:"stock" gorm:"type:int;not null"`
	Price       int    `json:"price" gorm:"type:int;not null"`
	UserId      string `json:"user_id" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
