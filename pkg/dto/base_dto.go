package dto

import (
	"time"
)

type BaseDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
