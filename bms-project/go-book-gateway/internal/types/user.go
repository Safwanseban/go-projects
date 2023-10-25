package types

import "github.com/google/uuid"

type User struct {
	ID          uint      `json:"id" `
	UUID        uuid.UUID `json:"uuid"`
	UserName    string    `json:"userName" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
}
