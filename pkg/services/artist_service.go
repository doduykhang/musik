package services

import (
	"fmt"

	"github.com/doduykhang/musik/pkg/constant"
	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
)

type ArtistService interface {
	CreateArtist(*dto.CreateArtistRequest) (*dto.ArtistDTO, error)
	UpdateArtist(*dto.UpdateArtistRequest) (*dto.ArtistDTO, error)
	DeleteArtist(uint) (*dto.ArtistDTO, error)
	FindArtists() (*[]dto.ArtistDTO, error)
	FindArtist(uint) (*dto.ArtistDTO, error)
}

type artistServiceImpl struct{}

func GetAritstServive() ArtistService {
	return &artistServiceImpl{}
}

func (service *artistServiceImpl) CreateArtist(request *dto.CreateArtistRequest) (*dto.ArtistDTO, error) {
	var artist models.Artist
	err := Map(&artist, request)
	if err != nil {
		return nil, err
	}

	err = uploadArtistImage(&request.ImageFile, &artist)
	if err != nil {
		return nil, err
	}
	result := db.Create(&artist)
	if result.Error != nil {
		return nil, result.Error
	}

	var artistDTO dto.ArtistDTO
	Map(&artistDTO, &artist)

	return &artistDTO, nil
}

func (service *artistServiceImpl) UpdateArtist(request *dto.UpdateArtistRequest) (*dto.ArtistDTO, error) {
	var artist models.Artist

	result := db.First(&artist, request.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	err := Map(&artist, request)
	if err != nil {
		return nil, err
	}

	updateResult := db.Save(&artist)
	if updateResult.Error != nil {
		return nil, result.Error
	}

	var artistDTO dto.ArtistDTO
	Map(&artistDTO, &artist)

	return &artistDTO, nil
}

func (service *artistServiceImpl) DeleteArtist(ID uint) (*dto.ArtistDTO, error) {
	var artist models.Artist
	artist.ID = ID
	result := db.First(&artist)
	if result.Error != nil {
		return nil, result.Error
	}
	db.Delete(&artist)
	var artistDTO dto.ArtistDTO
	err := Map(&artistDTO, &artist)
	if err != nil {
		return nil, err
	}
	return &artistDTO, nil
}

func (service *artistServiceImpl) FindArtists() (*[]dto.ArtistDTO, error) {
	var artists []models.Artist
	result := db.Find(&artists)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	var artistDTOs []dto.ArtistDTO
	err := Map(&artistDTOs, &artists)
	if err != nil {
		return nil, err
	}
	return &artistDTOs, nil
}

func (service *artistServiceImpl) FindArtist(ID uint) (*dto.ArtistDTO, error) {
	var artist models.Artist
	artist.ID = ID
	result := db.Find(&artist)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	var artistDTO dto.ArtistDTO

	err := Map(&artistDTO, &artist)
	if err != nil {
		return nil, err
	}
	return &artistDTO, nil
}

func uploadArtistImage(image *dto.MultipartForm, artist *models.Artist) error {
	_, _, err := fileService.SaveFile(image.Bytes, constant.ARTIST_IMAGE_PATH, image.Name)

	if err != nil {
		return err
	}

	artist.Image = constant.AUDIO_PATH + image.Name

	return nil
}
