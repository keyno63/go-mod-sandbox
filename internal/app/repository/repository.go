package repository

import (
	"database/sql"
	"go-mod2/internal/app/model"
)

/**
Repository
*/
type UserRepository interface {
	GetUser(id string) model.UserAccount
}

type UserRepositoryImpl struct {
	DbConnector *sql.DB
}

func (us UserRepositoryImpl) GetUser(id string) model.UserAccount {
	// 仮
	// TODO: DBとの接続の実装
	return model.UserAccount{
		Id:        id,
		FirstName: "first",
		LastName:  "last",
	}
}
