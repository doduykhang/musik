package dto

import (
	"time"
)

type BaseDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type MessageDTO struct {
	Message string `json:"message"`
}

type MultipartForm struct {
	Bytes []byte
	Name  string
}
