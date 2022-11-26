package dto

import (
	"time"
)

type BaseSongDTO struct {
	Title       string    `schema:"title" json:"title" validate:"required"`
	Duration    uint      `schema:"duration" json:"duration" validate:"required"`
	ReleaseDate time.Time `schema:"releaseDate" json:"releaseDate" validate:"required"`
}

type CreateSongRequest struct {
	BaseSongDTO
	Aritsts   []uint `schema:"artists" json:"artists" validate:"required"`
	CoverFile MultipartForm
	AudioFile MultipartForm
}

type UpdateSongRequest struct {
	BaseSongDTO
	ID      uint   `json:"id" validate:"required"`
	Aritsts []uint `schema:"artists" json:"artists" validate:"required"`
}

type UpdateAudioRequet struct {
	ID        uint `schema:"id"`
	AudioFile MultipartForm
}

type UpdateCoverRequet struct {
	ID        uint `schema:"id"`
	CoverFile MultipartForm
}

type SongDTO struct {
	BaseDTO
	BaseSongDTO
	Artists []ArtistDTO `json:"artists"`
	Cover   string      `json:"cover"`
	Audio   string      `json:"audio"`
}

type FindSongRequest struct {
	Pagination
	Title string `schema:"title"`
}
