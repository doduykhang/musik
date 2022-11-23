package models

type Album struct {
	BaseModel
	Name  string `gorm:"not null"`
	Image string `gorm:"not null"`
	Songs []Song `gorm:"many2many:albums_songs"`
}

func init() {
	db.SetupJoinTable(&Album{}, "Songs", &AlbumSong{})
	db.AutoMigrate(&Album{})
}
