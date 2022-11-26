package services

import (
	"github.com/doduykhang/musik/pkg/constant"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
	"github.com/doduykhang/musik/pkg/utils"
)

type SongService interface {
	CreateSong(*dto.CreateSongRequest) (*dto.SongDTO, error)
	UpdateSong(*dto.UpdateSongRequest) (*dto.SongDTO, error)
	UpdateSongAudio(*dto.UpdateAudioRequet) (*dto.SongDTO, error)
	UpdateSongCover(*dto.UpdateCoverRequet) (*dto.SongDTO, error)
	DeleteSong(uint) (*dto.SongDTO, error)
	FindSongs(*dto.FindSongRequest) (*[]dto.SongDTO, error)
	FindSong(uint) (*dto.SongDTO, error)
}

type songServiceImpl struct{}

var (
	fileService FileService
)

func init() {
	fileService = GetLocalFileService()
}

func GetSongService() SongService {
	return &songServiceImpl{}
}

func (service *songServiceImpl) CreateSong(request *dto.CreateSongRequest) (*dto.SongDTO, error) {
	var song models.Song
	Map(&song, request)

	var artists []models.Artist
	result := db.Find(&artists, request.Aritsts)

	if result.Error != nil {
		return nil, result.Error
	}

	song.Artists = artists

	err := uploadSongAssets(
		&request.AudioFile,
		&request.CoverFile,
		&song,
	)

	if err != nil {
		return nil, result.Error
	}

	result = db.Save(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	var artistDTOs []dto.ArtistDTO
	Map(&artistDTOs, &artists)

	return &dto.SongDTO{
		BaseSongDTO: request.BaseSongDTO,
		Cover:       song.Cover,
		Audio:       song.Audio,
		BaseDTO: dto.BaseDTO{
			ID:        song.ID,
			CreatedAt: song.CreatedAt,
		},
		Artists: artistDTOs,
	}, nil
}

func (service *songServiceImpl) UpdateSong(request *dto.UpdateSongRequest) (*dto.SongDTO, error) {
	var song models.Song

	result := db.Find(&song, request.ID)

	if result.Error != nil {
		return nil, result.Error
	}

	Map(&song, request)

	var artists []models.Artist
	result = db.Find(&artists, request.Aritsts)

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
	Map(&artistDTOs, &artists)

	return &dto.SongDTO{
		BaseSongDTO: request.BaseSongDTO,
		BaseDTO: dto.BaseDTO{
			ID:        song.ID,
			CreatedAt: song.CreatedAt,
		},
		Artists: artistDTOs,
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
	Map(&songDTO, &song)
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
	Map(&songDTOs, &songs)
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
	Map(&songDTO, &song)
	return &songDTO, nil
}

func (service *songServiceImpl) UpdateSongAudio(request *dto.UpdateAudioRequet) (*dto.SongDTO, error) {
	var song models.Song
	result := db.Find(&song, request.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	err := uploadAudio(
		&request.AudioFile,
		&song,
	)
	if err != nil {
		return nil, result.Error
	}

	result = db.Save(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	var songDTO dto.SongDTO
	Map(&songDTO, &song)

	return &songDTO, nil
}
func (service *songServiceImpl) UpdateSongCover(request *dto.UpdateCoverRequet) (*dto.SongDTO, error) {
	var song models.Song
	result := db.Find(&song, request.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	err := uploadCover(
		&request.CoverFile,
		&song,
	)
	if err != nil {
		return nil, result.Error
	}

	result = db.Save(&song)
	if result.Error != nil {
		return nil, result.Error
	}

	var songDTO dto.SongDTO
	Map(&songDTO, &song)

	return &songDTO, nil

}

func uploadSongAssets(audio *dto.MultipartForm, cover *dto.MultipartForm, song *models.Song) error {
	err := uploadAudio(audio, song)
	if err != nil {
		return err
	}
	err = uploadCover(cover, song)
	return err
}

func uploadAudio(audio *dto.MultipartForm, song *models.Song) error {
	_, _, err := fileService.SaveFile(audio.Bytes, constant.AUDIO_PATH, audio.Name)

	if err != nil {
		return err
	}

	song.Audio = constant.AUDIO_PATH + audio.Name

	return nil
}

func uploadCover(cover *dto.MultipartForm, song *models.Song) error {
	_, _, err := fileService.SaveFile(cover.Bytes, constant.COVER_PATH, cover.Name)

	if err != nil {
		return err
	}

	song.Cover = constant.COVER_PATH + cover.Name

	return nil
}
