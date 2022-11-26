package dto

type CreateArtistRequest struct {
	Name        string `schema:"name" validate:"required"`
	ImageFile   MultipartForm
	Description string `schema:"description" validate:"required"`
	BirthDate   string `schema:"birthDate" validate:"required"`
}

type UpdateArtistRequest struct {
	Name        string `schema:"name" validate:"required"`
	Description string `schema:"description" validate:"required"`
	BirthDate   string `schema:"birthDate" validate:"required"`
	ID          uint   `json:"id"`
}

type ArtistDTO struct {
	BaseDTO
	Name        string `json:"name" validate:"required"`
	Image       string `json:"image" validate:"required"`
	Description string `json:"description" validate:"required"`
	BirthDate   string `json:"birthDate" validate:"required"`
}
