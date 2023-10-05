package types

import (
	"time"

	"github.com/google/uuid"
)

type Places struct {
	ID           uint `json:"id"`
	UUID         uuid.UUID
	LocationName string
	Country      string
	State        string
	Theators     []Theator
}

type Theator struct {
	ID          uint `json:"id"`
	UUID        uuid.UUID
	TheatorName string
	Screens     []Screen
}
type Screen struct {
	ID           uint `json:"id"`
	UUID         uuid.UUID
	ScreenNumber string
	AudioType    string
	ScreenType   string
	Tiers        []Tier
	CurrentMovie Movie
}
type Tier struct {
	ID          uint `json:"id"`
	UUID        uuid.UUID
	TierName    string
	ListOfSeats []Seats
	TierPrice   uint
}
type Seats struct {
	ID         uint `json:"id"`
	UUID       uuid.UUID
	TypeOfSeat string
}

type Movie struct {
	ID                 uint `json:"id"`
	UUID               uuid.UUID
	MovieName          string
	NativeLanguage     Language
	AvailableLanguages []Language
	Genre              string
}

type Ticket struct {
	ID        uint `json:"id"`
	UUID      uuid.UUID
	Price     uint
	Seats     []Seats
	Tax       uint
	StartTime time.Time
	EndTime   time.Time
	GivenTime time.Time
}
type Language struct {
	ID   uint `json:"id"`
	UUID uuid.UUID
	Name string
}
type User struct {
	ID          uint `json:"id"`
	UUID        uuid.UUID
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
}
