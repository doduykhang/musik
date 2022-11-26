package dto

import (
	"time"
)

type BaseSongDTO struct {
	Title       string    `json:"title" validate:"required"`
	Cover       string    `json:"cover" validate:"required"`
	Duration    uint      `json:"duration" validate:"required"`
	ReleaseDate time.Time `json:"releaseDate" validate:"required"`
	Audio       string    `json:"audio" validate:"required"`
}

type CreateSongRequest struct {
	BaseSongDTO
	Aritsts []uint `json:"artists" validate:"required"`
}

type UpdateSongRequest struct {
	CreateSongRequest
	ID uint `json:"id" validate:"required"`
}

type SongDTO struct {
	BaseDTO
	BaseSongDTO
	ArtistDTOs []ArtistDTO `json:"artists"`
}

type FindSongRequest struct {
	Pagination
	Title string `schema:"title"`
}
