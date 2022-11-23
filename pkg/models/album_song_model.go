package models

type AlbumSong struct {
	SongID      uint `gorm:"primaryKey"`
	AlbumID     uint `gorm:"primaryKey"`
	TrackNumber uint `gorm:"not null"`
}
