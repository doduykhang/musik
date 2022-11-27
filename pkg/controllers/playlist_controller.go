package controllers

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/services"
	"github.com/doduykhang/musik/pkg/utils"
)

var (
	playlistService services.PlaylistService
)

func init() {
	playlistService = services.GetPlaylistService()
}

func CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	var request dto.CreatePlaylistRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	err = validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	createdArtist, err := playlistService.CreatePlaylist(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdArtist)
}

func UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdatePlaylistRequest
	utils.ParseBody(r, &request)

	err := validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	updatedArtist, err := playlistService.UpdatePlaylist(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, updatedArtist)
}

func DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	deletedArtist, err := playlistService.DeletePlaylist(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, deletedArtist)
}

func FindPlaylists(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	playlists, err := playlistService.FindPlaylists(ID)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, playlists)
}

func FindPlaylist(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	artist, err := playlistService.FindPlaylist(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, artist)
}

func AddSongToPlaylist(w http.ResponseWriter, r *http.Request) {
	var request dto.AddSongToPlaylistRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}

	err = validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	artist, err := playlistService.AddSongToPlaylist(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, artist)
}

func RemoveSongFromPlaylist(w http.ResponseWriter, r *http.Request) {
	var request dto.RemoveSongFromPlaylistRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}

	err = validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	artist, err := playlistService.RemoveSongFromPlaylist(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, artist)
}
