package models

import (
	"time"
)

type Song struct {
	BaseModel
	Title       string     `gorm:"not null" json:"title"`
	Cover       string     `gorm:"not null" json:"cover"`
	Duration    uint       `gorm:"not null" json:"duration"`
	ReleaseDate time.Time  `gorm:"not null" json:"releaseDate"`
	Audio       string     `gorm:"not null" json:"audio"`
	Playlists   []Playlist `gorm:"many2many:songs_playlists"`
	Artists     []Artist   `gorm:"many2many:songs_artists"`
}

func init() {
	db.AutoMigrate(&Song{})
}
