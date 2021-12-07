//go:generate mockgen -source=repository.go -destination=../../../mock/mock_repository.go

package repository

import (
	"database/sql"
	"go-mod-sandbox/internal/app/domain/model"
	"go-mod-sandbox/internal/app/domain/model/dto"
	"strconv"
)

/*
UserRepository の実装です
*/
type UserRepository interface {
	GetUser(id string) (*model.UserAccount, error)
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return UserRepositoryImpl{
		dbConnector: db,
	}
}

type UserRepositoryImpl struct {
	dbConnector *sql.DB
}

func (us UserRepositoryImpl) GetUser(id string) (*model.UserAccount, error) {

	cmd := "SELECT * FROM user WHERE $1"
	rows, err := us.dbConnector.Query(cmd, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ua dto.UserDto
	if err := rows.Scan(&ua.ID, &ua.FirstName, &ua.LastName); err == nil {
		return nil, err
	}

	i := strconv.Itoa(ua.ID)
	return &model.UserAccount{
		ID:        i,
		FirstName: ua.FirstName,
		LastName:  ua.LastName,
	}, nil
}
