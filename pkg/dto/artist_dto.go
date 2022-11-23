package dto

type CreateArtistRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	BirthDate   string `json:"birthDate"`
}

type UpdateArtistRequest struct {
	CreateArtistRequest
	ID uint `json:"id"`
}

type ArtistDTO struct {
	BaseDTO
	CreateArtistRequest
}
