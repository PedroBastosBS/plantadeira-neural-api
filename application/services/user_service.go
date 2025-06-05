package services

import (
	"errors"
	"github.com/google/uuid"
	"plantadeira-neural-api/domain"
	"plantadeira-neural-api/framework/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Save(user *domain.User) error {
	if user == nil {
		return errors.New("Usuario não informado!")
	}
	user.ID = uuid.New().String()
	if err := user.Validate(); err != nil {
		return err
	}

	return s.userRepo.Save(user)
}
