package dto

type PlaylistDTO struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CreatePlaylistRequest struct {
	Name   string `json:"name" validate:"required"`
	UserID uint   `json:"userID" validate:"required"`
}

type UpdatePlaylistRequest struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type AddSongToPlaylistRequest struct {
	SongID     uint `json:"songID" validate:"required"`
	PlaylistID uint `json:"playlistID" validate:"required"`
}

type RemoveSongFromPlaylistRequest struct {
	SongID     uint `json:"songID" validate:"required"`
	PlaylistID uint `json:"playlistID" validate:"required"`
}
