package controllers

import (
	"fmt"
	"net/http"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/services"
	"github.com/doduykhang/musik/pkg/utils"
)

var (
	artistService services.ArtistService
)

func init() {
	artistService = services.GetAritstServive()
}

func CreateArtist(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateArtistRequest

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		fmt.Println("test1")
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	err = decoder.Decode(&request, r.PostForm)

	if err != nil {
		fmt.Println("test2")
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	//get audio
	imageByte, imageName, err := utils.GetFileByteWithName(r, "image")

	request.ImageFile.Name = imageName
	request.ImageFile.Bytes = imageByte

	err = validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	createdArtist, err := artistService.CreateArtist(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdArtist)
}

func UpdateArtist(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateArtistRequest
	utils.ParseBody(r, &request)

	err := validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	updatedArtist, err := artistService.UpdateArtist(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, updatedArtist)
}

func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	deletedArtist, err := artistService.DeleteArtist(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, deletedArtist)
}

func FindAllArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := artistService.FindArtists()
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
	artist, err := artistService.FindArtist(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, artist)
}
