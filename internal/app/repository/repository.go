package repository

import (
	"database/sql"
	"go-mod2/internal/app/model/dto"
)

/*
UserRepository の実装です
*/
type UserRepository interface {
	GetUser(id string) dto.UserDto
}

func NewUserRepositoryImpl(db *sql.DB) UserRepository {
	return UserRepositoryImpl{
		dbConnector: db,
	}
}

type UserRepositoryImpl struct {
	dbConnector *sql.DB
}

func (us UserRepositoryImpl) GetUser(id string) dto.UserDto {
	var ua dto.UserDto

	cmd := "SELECT * FROM user WHERE $1"
	rows, err := us.dbConnector.Query(cmd, id)
	if err != nil {
		return ua
	}
	defer rows.Close()
	if err := rows.Scan(&ua.Id, &ua.FirstName, &ua.LastName); err != nil {
		return ua
	}
	return ua
}
