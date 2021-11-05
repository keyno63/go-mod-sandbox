//go:generate mockgen -source=service.go -destination=../../../mock/mock_service.go

package service

import (
	"go-mod-sandbox/internal/app/model"
	"go-mod-sandbox/internal/app/repository"
)

/*
UserService の実装です
*/
type UserService interface {
	GetUser(id string) (*model.UserAccount, error)
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
func (us UserServiceImpl) GetUser(id string) (*model.UserAccount, error) {
	ua, err := us.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return ua, nil
}
