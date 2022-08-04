package auth

import "project/online-store/entity"

type AuthService interface {
	Register(user entity.User) error
}
