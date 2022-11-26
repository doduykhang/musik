package dto

type CreateArtistRequest struct {
	Name        string `json:"name" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Description string `json:"description" validate:"required"`
	BirthDate   string `json:"birthDate" validate:"required"`
}

type UpdateArtistRequest struct {
	CreateArtistRequest
	ID uint `json:"id"`
}

type ArtistDTO struct {
	BaseDTO
	CreateArtistRequest
}
