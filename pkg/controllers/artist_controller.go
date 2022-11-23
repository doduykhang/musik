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
	createdArtist, _ := services.CreateArtist(&request)
	utils.JsonResponse(&w, createdArtist)
}

func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateArtistRequest
	utils.ParseBody(r, &request)
	updatedArtist, _ := services.UpdateArtist(&request)
	utils.JsonResponse(&w, updatedArtist)
}

func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	ID, _ := utils.GetIDFromRequest(r)

	deletedArtist, _ := services.DeleteAritst(ID)
	utils.JsonResponse(&w, deletedArtist)
}

func FindAllArtists(w http.ResponseWriter, r *http.Request) {
	artists, _ := services.FindAllArtist()
	utils.JsonResponse(&w, artists)
}

func FindArtist(w http.ResponseWriter, r *http.Request) {
	ID, _ := utils.GetIDFromRequest(r)
	artist, _ := services.FindArtist(ID)
	utils.JsonResponse(&w, artist)
}
