package types

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint
	UUID        uuid.UUID
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
}



func (u *User) HashPassword() error {

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("password hashing error")
	}
	u.Password = string(password)
	return nil

}
