package controller

import (
	"go-mod2/internal/app/model"
	"go-mod2/internal/app/service"
)

/**
Controller
*/
type UserController interface {
	GetUser(id string) model.UserAccount
}

func NewUserController(service service.UserService) UserController {
	return UserControllerImpl{
		service,
	}
}

type UserControllerImpl struct {
	userService service.UserService
}

func (uc UserControllerImpl) GetUser(id string) model.UserAccount {
	return uc.userService.GetUser(id)
}
