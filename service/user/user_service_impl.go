package user

import (
	"project/online-store/config"
	"project/online-store/entity"
	"project/online-store/repository/user"
)

type Service struct {
	r    user.UserRepository
	conf config.Config
}

func NewUserService(r user.UserRepository, conf config.Config) UserService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) FindAllUsers() ([]entity.User, error) {
	user, err := s.r.FindAllUsers()
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *Service) FindUserById(id string) (entity.User, error) {
	user, err := s.r.FindUserById(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
