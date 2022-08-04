package auth

import (
	"project/online-store/config"
	"project/online-store/entity"
	"project/online-store/repository/auth"
)

type Service struct {
	r    auth.AuthRepository
	conf config.Config
}

func NewAuthService(r auth.AuthRepository, conf config.Config) AuthService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) Register(user entity.User) error {
	err := s.r.Register(user)
	if err != nil {
		return err
	}

	return nil
}
