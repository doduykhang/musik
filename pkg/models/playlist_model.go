package models

type Playlist struct {
	BaseModel
	Name   string `gorm:"not null"`
	Image  string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	Songs  []Song `gorm:"many2many:songs_playlists"`
}

func init() {
	db.AutoMigrate(&Playlist{})
}
