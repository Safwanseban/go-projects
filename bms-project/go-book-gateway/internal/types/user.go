package types

import "github.com/google/uuid"

type User struct {
	ID          uint `json:"id"`
	UUID        uuid.UUID
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
}
