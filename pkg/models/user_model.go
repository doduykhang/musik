package models

import (
	"time"
)

type User struct {
	BaseModel
	Firstname string    `gorm:"not null"`
	Lastname  string    `gorm:"not null"`
	BirthDate time.Time `gorm:"not null"`
	Gender    bool      `gorm:"not null"`
	Image     string    `gorm:"not null"`
	AccountID int       `gorm:"not null"`
	Account   Account
	Playlists []Playlist
}

func init() {
	db.AutoMigrate(&User{})
}
