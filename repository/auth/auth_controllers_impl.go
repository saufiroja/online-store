package auth

import (
	"project/online-store/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Register(user entity.User) error {
	err := r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
