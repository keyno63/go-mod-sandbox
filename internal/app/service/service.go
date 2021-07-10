package service

import (
	"go-mod2/internal/app/model"
	"go-mod2/internal/app/repository"
	"strconv"
)

/*
UserService の実装です
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

// GetUser は DTO を model に変換までします
func (us UserServiceImpl) GetUser(id string) model.UserAccount {
	ua := us.userRepository.GetUser(id)
	i := strconv.Itoa(ua.Id)
	return model.UserAccount{
		Id:        i,
		FirstName: ua.FirstName,
		LastName:  ua.LastName,
	}
}
