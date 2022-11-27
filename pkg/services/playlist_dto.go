package services

import (
	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
)

type PlaylistService interface {
	CreatePlaylist(request *dto.CreatePlaylistRequest) (*dto.PlaylistDTO, error)
	UpdatePlaylist(request *dto.UpdatePlaylistRequest) (*dto.PlaylistDTO, error)
	DeletePlaylist(id uint) (*dto.PlaylistDTO, error)
	FindPlaylists(userID uint) (*[]dto.PlaylistDTO, error)
	FindPlaylist(id uint) (*dto.PlaylistDTO, error)
	AddSongToPlaylist(request *dto.AddSongToPlaylistRequest) (*dto.PlaylistDTO, error)
	RemoveSongFromPlaylist(request *dto.RemoveSongFromPlaylistRequest) (*dto.PlaylistDTO, error)
}

type playlistServiceImpl struct{}

func (service *playlistServiceImpl) CreatePlaylist(request *dto.CreatePlaylistRequest) (*dto.PlaylistDTO, error) {
	var playlist models.Playlist
	Map(&playlist, request)
	result := db.Create(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}
	var playlistDTO dto.PlaylistDTO
	Map(&playlistDTO, &playlist)
	return &playlistDTO, nil

}
func (service *playlistServiceImpl) UpdatePlaylist(request *dto.UpdatePlaylistRequest) (*dto.PlaylistDTO, error) {
	var playlist models.Playlist
	result := db.Find(&playlist, request.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	Map(&playlist, request)
	result = db.Save(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}

	var playlistDTO dto.PlaylistDTO
	Map(&playlistDTO, &playlist)
	return &playlistDTO, nil
}
func (service *playlistServiceImpl) DeletePlaylist(id uint) (*dto.PlaylistDTO, error) {
	var playlist models.Playlist
	result := db.Find(&playlist, id)
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.Delete(&playlist)
	if result.Error != nil {
		return nil, result.Error
	}

	var playlistDTO dto.PlaylistDTO
	Map(&playlistDTO, &playlist)
	return &playlistDTO, nil
}
func (service *playlistServiceImpl) FindPlaylists(userID uint) (*[]dto.PlaylistDTO, error) {
	var user models.User
	result := db.Preload("Playlists").Find(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	var playlistDTOs []dto.PlaylistDTO
	Map(&playlistDTOs, &user.Playlists)
	return &playlistDTOs, nil
}
func (service *playlistServiceImpl) FindPlaylist(id uint) (*dto.PlaylistDTO, error) {
	var playlist models.Playlist
	result := db.Find(&playlist, id)
	if result.Error != nil {
		return nil, result.Error
	}
	var playlistDTO dto.PlaylistDTO
	Map(&playlistDTO, &playlist)
	return &playlistDTO, nil
}
func (service *playlistServiceImpl) AddSongToPlaylist(request *dto.AddSongToPlaylistRequest) (*dto.PlaylistDTO, error) {
	var playlist models.Playlist
	result := db.Find(&playlist, request.PlaylistID)
	if result.Error != nil {
		return nil, result.Error
	}

	var song models.Song
	result = db.Find(&song, request.SongID)
	if result.Error != nil {
		return nil, result.Error
	}

	db.Model(&playlist).Association("Songs").Append(&song)

	var playlistDTO dto.PlaylistDTO
	Map(&playlistDTO, &playlist)
	return &playlistDTO, nil
}
func (service *playlistServiceImpl) RemoveSongFromPlaylist(request *dto.RemoveSongFromPlaylistRequest) (*dto.PlaylistDTO, error) {
	var playlist models.Playlist
	result := db.Find(&playlist, request.PlaylistID)
	if result.Error != nil {
		return nil, result.Error
	}

	var song models.Song
	result = db.Find(&song, request.SongID)
	if result.Error != nil {
		return nil, result.Error
	}

	db.Model(&playlist).Association("Songs").Delete(song)

	var playlistDTO dto.PlaylistDTO
	Map(&playlistDTO, &playlist)
	return &playlistDTO, nil
}
