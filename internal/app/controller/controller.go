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

type UserControllerImpl struct {
	UserService service.UserService
}

func (uc UserControllerImpl) GetUser(id string) model.UserAccount {
	return uc.UserService.GetUser(id)
}
