package dto

import (
	"time"
)

type CreateSongRequest struct {
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	Duration    uint      `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`
	Audio       string    `json:"audio"`
}

type SongDTO struct {
	BaseDTO
	CreateSongRequest
}
