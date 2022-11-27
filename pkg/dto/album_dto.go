package dto

type AlbumDTO struct {
	BaseDTO
	Name  string `json:"name"`
	Image string `json:"image"`
}

type CreateAlbumRequest struct {
	Name      string `schema:"name" validate:"required"`
	ImageFile MultipartForm
}

type UpdateAlbumRequest struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateAlbumImageRequest struct {
	ID        uint `json:"id" validate:"required"`
	ImageFile MultipartForm
}

type AddSongToAlbumRequest struct {
	SongIDs []uint `json:"songIDs" validate:"required"`
	AlbumID uint   `json:"albumID" validate:"required"`
}

type RemoveSongFromAlbumRequest struct {
	SongIDs []uint `json:"songIDs" validate:"required"`
	AlbumID uint   `json:"albumID" validate:"required"`
}

type FindAlbumRequest struct {
	Pagination
	Name string `schema:"name" validate:"required"`
}
