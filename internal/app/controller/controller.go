package controller

import (
	"go-mod2/internal/app/model"
	"go-mod2/internal/app/service"
)

/*
UserController の実装です
*/
type UserController interface {
	GetUser(id string) (*model.UserAccount, error)
}

func NewUserController(service service.UserService) UserController {
	return UserControllerImpl{
		service,
	}
}

type UserControllerImpl struct {
	userService service.UserService
}

func (uc UserControllerImpl) GetUser(id string) (*model.UserAccount, error) {
	return uc.userService.GetUser(id)
}
