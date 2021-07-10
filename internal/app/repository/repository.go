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

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return UserRepositoryImpl{
		dbConnector: db,
	}
}

type UserRepositoryImpl struct {
	dbConnector *sql.DB
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
