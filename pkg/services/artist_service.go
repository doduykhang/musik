package services

import (
	"fmt"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
	"github.com/doduykhang/musik/pkg/utils"
)

func CreateArtist(request *dto.CreateArtistRequest) (*dto.ArtistDTO, error) {
	var artist models.Artist
	utils.ConverseStruct(request, &artist)
	result := db.Create(&artist)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.ArtistDTO{
		CreateArtistRequest: *request,
		BaseDTO: dto.BaseDTO{
			ID:        artist.ID,
			CreatedAt: artist.CreatedAt,
		},
	}, nil
}

func UpdateArtist(request *dto.UpdateArtistRequest) (*dto.ArtistDTO, error) {
	var artist models.Artist
	utils.ConverseStruct(request, &artist)
	result := db.Save(&artist)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.ArtistDTO{
		CreateArtistRequest: *&request.CreateArtistRequest,
		BaseDTO: dto.BaseDTO{
			ID:        artist.ID,
			CreatedAt: artist.CreatedAt,
		},
	}, nil
}

func DeleteAritst(ID uint) (*dto.ArtistDTO, error) {
	var artist models.Artist
	artist.ID = ID
	result := db.First(&artist)
	if result.Error != nil {
		return nil, result.Error
	}
	db.Delete(&artist)
	var artistDTO dto.ArtistDTO
	utils.ConverseStruct(&artist, &artistDTO)
	return &artistDTO, nil
}

func FindAllArtist() (*[]dto.ArtistDTO, error) {
	var artists []models.Artist
	result := db.Find(&artists)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	var artistDTOs []dto.ArtistDTO
	utils.ConverseStruct(&artists, &artistDTOs)
	return &artistDTOs, nil
}

func FindArtist(ID uint) (*dto.ArtistDTO, error) {
	var artist models.Artist
	artist.ID = ID
	result := db.Find(&artist)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	var artistDTO dto.ArtistDTO
	utils.ConverseStruct(&artist, &artistDTO)
	return &artistDTO, nil
}
