package controllers

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/services"
	"github.com/doduykhang/musik/pkg/utils"
)

var (
	albumService services.AlbumService
)

func init() {
	albumService = services.GetAlbumService()
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateAlbumRequest

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	err = decoder.Decode(&request, r.PostForm)

	if err != nil {
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

	createdAlbum, err := albumService.CreateAlbum(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdAlbum)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateAlbumRequest
	utils.ParseBody(r, &request)

	err := validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	updatedAlbum, err := albumService.UpdateAlbum(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, updatedAlbum)
}

func UpdateAlbumImage(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateAlbumImageRequest

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	err = decoder.Decode(&request, r.PostForm)

	if err != nil {
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

	createdAlbum, err := albumService.UpdateAlbumImage(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdAlbum)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	dto, err := albumService.DeleteAlbum(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, dto)
}

func FindAllAlbums(w http.ResponseWriter, r *http.Request) {
	var request dto.FindAlbumRequest
	err := decoder.Decode(&request, r.URL.Query())
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}

	dtos, err := albumService.FindAlbums(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, dtos)
}

func FindAlbum(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	dto, err := albumService.FindAlbum(ID)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, dto)
}

func AddSongs(w http.ResponseWriter, r *http.Request) {
	var request []dto.AddSongToAlbumRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	dto, err := albumService.AddSong(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, dto)
}

func RemoveSongs(w http.ResponseWriter, r *http.Request) {
	var request []dto.RemoveSongFromAlbumRequest
	err := utils.ParseBody(r, &request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}
	dto, err := albumService.RemoveSong(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, dto)
}
