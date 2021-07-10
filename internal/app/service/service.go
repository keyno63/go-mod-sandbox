package service

import (
	"go-mod2/internal/app/model"
	"go-mod2/internal/app/repository"
)

/**
Service
*/
type UserService interface {
	GetUser(id string) model.UserAccount
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return UserServiceImpl{
		repository,
	}
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func (us UserServiceImpl) GetUser(id string) model.UserAccount {
	return us.userRepository.GetUser(id)
}
