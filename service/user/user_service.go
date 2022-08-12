package user

import "project/online-store/entity"

type UserService interface {
	FindAllUsers() ([]entity.User, error)
	FindUserById(id string) (entity.User, error)
}
