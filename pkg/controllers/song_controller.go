package controllers

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/services"
	"github.com/doduykhang/musik/pkg/utils"
)

var (
	songService services.SongService
)

func init() {
	songService = services.GetSongService()
}

func CreateSong(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateSongRequest
	utils.ParseBody(r, &request)

	err := validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	createdSong, err := songService.CreateSong(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdSong)
}

func UpdateSong(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateSongRequest
	utils.ParseBody(r, &request)

	err := validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	updatedSong, err := songService.UpdateSong(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, updatedSong)
}

func DeleteSong(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	deletedSong, err := songService.DeleteSong(ID)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, deletedSong)
}

func FindSong(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	deletedSong, err := songService.FindSong(ID)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, deletedSong)
}

func FindSongs(w http.ResponseWriter, r *http.Request) {
	var request dto.FindSongRequest
	err := decoder.Decode(&request, r.URL.Query())

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	err = validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	songs, err := songService.FindSongs(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, songs)
}
