package user

import "project/online-store/entity"

type UserRepository interface {
	FindAllUsers() ([]entity.User, error)
}
