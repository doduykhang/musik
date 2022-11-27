package services

import (
	"fmt"

	"github.com/doduykhang/musik/pkg/constant"
	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
	"github.com/doduykhang/musik/pkg/utils"
)

type AlbumService interface {
	CreateAlbum(*dto.CreateAlbumRequest) (*dto.AlbumDTO, error)
	UpdateAlbum(*dto.UpdateAlbumRequest) (*dto.AlbumDTO, error)
	UpdateAlbumImage(*dto.UpdateAlbumImageRequest) (*dto.AlbumDTO, error)
	DeleteAlbum(uint) (*dto.AlbumDTO, error)
	FindAlbums(*dto.FindAlbumRequest) (*[]dto.AlbumDTO, error)
	FindAlbum(uint) (*dto.AlbumDTO, error)
	AddSong(*[]dto.AddSongToAlbumRequest) (*[]dto.AddSongToAlbumRequest, error)
	RemoveSong(*[]dto.RemoveSongFromAlbumRequest) (*dto.AlbumDTO, error)
}

type albumServiceImpl struct {
}

func (service *albumServiceImpl) CreateAlbum(request *dto.CreateAlbumRequest) (*dto.AlbumDTO, error) {
	var album models.Album
	Map(&album, request)

	err := uploadAlbumImage(&request.ImageFile, &album)
	if err != nil {
		return nil, err
	}

	result := db.Create(&album)
	if result.Error != nil {
		return nil, result.Error
	}

	var albumDTO dto.AlbumDTO
	Map(&albumDTO, &album)
	return &albumDTO, nil
}
func (service *albumServiceImpl) UpdateAlbum(request *dto.UpdateAlbumRequest) (*dto.AlbumDTO, error) {
	var album models.Album
	result := db.Find(&album, request.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	Map(&album, request)
	result = db.Save(&album)
	if result.Error != nil {
		return nil, result.Error
	}

	var albumDTO dto.AlbumDTO
	Map(&albumDTO, album)
	return &albumDTO, nil

}
func (service *albumServiceImpl) UpdateAlbumImage(request *dto.UpdateAlbumImageRequest) (*dto.AlbumDTO, error) {
	var album models.Album
	result := db.Find(&album, request.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	err := uploadAlbumImage(&request.ImageFile, &album)
	if err != nil {
		return nil, err
	}

	result = db.Save(&album)
	if result.Error != nil {
		return nil, result.Error
	}

	var albumDTO dto.AlbumDTO
	Map(&albumDTO, album)
	return &albumDTO, nil
}
func (service *albumServiceImpl) DeleteAlbum(id uint) (*dto.AlbumDTO, error) {
	var album models.Album
	result := db.Find(&album, id)
	if result.Error != nil {
		return nil, result.Error
	}
	result = db.Delete(&album)
	if result.Error != nil {
		return nil, result.Error
	}
	var albumDTO dto.AlbumDTO
	Map(&albumDTO, &album)
	return &albumDTO, nil
}
func (service *albumServiceImpl) FindAlbums(request *dto.FindAlbumRequest) (*[]dto.AlbumDTO, error) {
	var albums []models.Album
	pagination := utils.Paginate(request.Page, request.Size)
	result := db.Scopes(pagination).Where("name LIKE ?", "%"+request.Name+"%").Find(&albums)
	if result.Error != nil {
		return nil, result.Error
	}

	var albumDTOs []dto.AlbumDTO
	Map(&albumDTOs, &albums)
	return &albumDTOs, nil
}
func (service *albumServiceImpl) FindAlbum(id uint) (*dto.AlbumDTO, error) {
	var album models.Album
	result := db.Find(&album, id)
	if result.Error != nil {
		return nil, result.Error
	}
	var albumDTO dto.AlbumDTO
	Map(&albumDTO, &album)
	return &albumDTO, nil
}
func (service *albumServiceImpl) AddSong(request *[]dto.AddSongToAlbumRequest) (*[]dto.AddSongToAlbumRequest, error) {
	var albums []models.AlbumSong
	Map(&albums, request)

	fmt.Println(albums)

	result := db.Create(&albums)
	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}
func (service *albumServiceImpl) RemoveSong(request *[]dto.RemoveSongFromAlbumRequest) (*dto.AlbumDTO, error) {
	var albums []models.AlbumSong
	Map(&albums, request)

	fmt.Println(albums)

	result := db.Delete(&albums)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.AlbumDTO{}, nil
}

func uploadAlbumImage(image *dto.MultipartForm, album *models.Album) error {
	_, _, err := fileService.SaveFile(image.Bytes, constant.ALBUM_IMAGE_PATH, image.Name)

	if err != nil {
		return err
	}

	album.Image = constant.ALBUM_IMAGE_PATH + image.Name

	return nil
}
