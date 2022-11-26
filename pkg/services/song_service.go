package services

import (
	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
	"github.com/doduykhang/musik/pkg/utils"
)

type SongService interface {
	CreateSong(*dto.CreateSongRequest) (*dto.SongDTO, error)
	UpdateSong(*dto.UpdateSongRequest) (*dto.SongDTO, error)
	DeleteSong(uint) (*dto.SongDTO, error)
	FindSongs(*dto.FindSongRequest) (*[]dto.SongDTO, error)
	FindSong(uint) (*dto.SongDTO, error)
}

type songServiceImpl struct{}

func GetSongService() SongService {
	return &songServiceImpl{}
}

func (service *songServiceImpl) CreateSong(request *dto.CreateSongRequest) (*dto.SongDTO, error) {
	var song models.Song
	utils.ConverseStruct(request, &song)

	var artists []models.Artist
	result := db.Find(&artists, request.Aritsts)

	if result.Error != nil {
		return nil, result.Error
	}

	song.Artists = artists
	result = db.Save(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	var artistDTOs []dto.ArtistDTO
	utils.ConverseStruct(&artists, &artistDTOs)

	return &dto.SongDTO{
		BaseSongDTO: request.BaseSongDTO,
		BaseDTO: dto.BaseDTO{
			ID:        song.ID,
			CreatedAt: song.CreatedAt,
		},
		ArtistDTOs: artistDTOs,
	}, nil
}

func (service *songServiceImpl) UpdateSong(request *dto.UpdateSongRequest) (*dto.SongDTO, error) {
	var song models.Song
	utils.ConverseStruct(request, &song)

	var artists []models.Artist
	result := db.Find(&artists, request.Aritsts)

	if result.Error != nil {
		return nil, result.Error
	}

	db.Model(&song).Association("Artists").Replace(artists)

	song.Artists = artists
	result = db.Save(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	var artistDTOs []dto.ArtistDTO
	utils.ConverseStruct(&artists, &artistDTOs)

	return &dto.SongDTO{
		BaseSongDTO: request.BaseSongDTO,
		BaseDTO: dto.BaseDTO{
			ID:        song.ID,
			CreatedAt: song.CreatedAt,
		},
		ArtistDTOs: artistDTOs,
	}, nil
}

func (service *songServiceImpl) DeleteSong(ID uint) (*dto.SongDTO, error) {
	var song models.Song

	result := db.Find(&song, ID)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Delete(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	var songDTO dto.SongDTO
	utils.ConverseStruct(&song, &songDTO)
	return &songDTO, nil
}

func (service *songServiceImpl) FindSongs(request *dto.FindSongRequest) (*[]dto.SongDTO, error) {
	var songs []models.Song
	pagination := utils.Paginate(request.Page, request.Size)
	result := db.Preload("Artists").Scopes(pagination).Where("title LIKE ?", "%"+request.Title+"%").Find(&songs)
	if result.Error != nil {
		return nil, result.Error
	}

	var songDTOs []dto.SongDTO
	utils.ConverseStruct(&songs, &songDTOs)
	return &songDTOs, nil
}

func (service *songServiceImpl) FindSong(ID uint) (*dto.SongDTO, error) {
	song := models.Song{}
	song.ID = ID
	result := db.Preload("Artists").Find(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	var songDTO dto.SongDTO
	utils.ConverseStruct(&song, &songDTO)
	return &songDTO, nil
}
