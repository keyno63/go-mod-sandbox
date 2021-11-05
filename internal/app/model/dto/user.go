package dto

import (
	"fmt"
)

type UserDto struct {
	ID        int
	FirstName string
	LastName  string
}

func (u UserDto) GoString() string {
	return fmt.Sprintf("gostring is: %s", "hoge")
}
