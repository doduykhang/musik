package controllers

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/services"
	"github.com/doduykhang/musik/pkg/utils"
)

var (
	songService services.SongService
	fileService services.FileService
)

func init() {
	songService = services.GetSongService()
	fileService = services.GetLocalFileService()
}

func CreateSong(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateSongRequest
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
	audioByte, audioName, err := utils.GetFileByteWithName(r, "audio")

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	request.AudioFile = dto.MultipartForm{Bytes: audioByte, Name: audioName}

	//get cover
	corverByte, corverName, err := utils.GetFileByteWithName(r, "cover")

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	request.CoverFile = dto.MultipartForm{Bytes: corverByte, Name: corverName}

	err = validate.Struct(&request)

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

func UpdateSongAudio(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateAudioRequet
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
	audioByte, audioName, err := utils.GetFileByteWithName(r, "audio")

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	request.AudioFile = dto.MultipartForm{Bytes: audioByte, Name: audioName}

	err = validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	createdSong, err := songService.UpdateSongAudio(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdSong)

}

func UpdateSongCover(w http.ResponseWriter, r *http.Request) {
	var request dto.UpdateCoverRequet
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
	audioByte, audioName, err := utils.GetFileByteWithName(r, "cover")

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	request.CoverFile = dto.MultipartForm{Bytes: audioByte, Name: audioName}

	err = validate.Struct(&request)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	createdSong, err := songService.UpdateSongCover(&request)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, createdSong)

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

func FindSongOfArtist(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	deletedSong, err := songService.FindSongOfArtist(ID)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, deletedSong)
}

func FindSongOfAlbum(w http.ResponseWriter, r *http.Request) {
	ID, err := utils.GetIDFromRequest(r)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 400)
		return
	}

	deletedSong, err := songService.FindSongOfAlbum(ID)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}
	utils.JsonResponse(&w, deletedSong)
}

func TestUpload(w http.ResponseWriter, r *http.Request) {
	bytes, handler, err := utils.GetFile(r, "audio")
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}

	_, _, err = fileService.SaveFile(bytes, "temp-folder", handler.Filename)

	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}

	utils.JsonResponse(&w, &dto.MessageDTO{
		Message: "ok",
	})
}
