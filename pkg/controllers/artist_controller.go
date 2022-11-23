package controllers

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/services"
	"github.com/doduykhang/musik/pkg/utils"
)

func CreateArtist(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateArtistRequest
	utils.ParseBody(r, &request)
	createdArtist, err := services.CreateArtist(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdArtist)
}

func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateArtistRequest
	utils.ParseBody(r, &request)
	updatedArtist, err := services.UpdateArtist(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, updatedArtist)
}

func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	deletedArtist, err := services.DeleteAritst(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, deletedArtist)
}

func FindAllArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := services.FindAllArtist()
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, artists)
}

func FindArtist(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	artist, err := services.FindArtist(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, artist)
}
