package routes

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterSongRoute(r *mux.Router) {
	r.HandleFunc("/song", controllers.CreateSong).Methods(http.MethodPost)
	r.HandleFunc("/song", controllers.UpdateSong).Methods(http.MethodPut)
	r.HandleFunc("/song", controllers.FindSongs).Methods(http.MethodGet)
	r.HandleFunc("/song/artist/{ID}", controllers.FindSongOfArtist).Methods(http.MethodGet)
	r.HandleFunc("/song/album/{ID}", controllers.FindSongOfAlbum).Methods(http.MethodGet)
	r.HandleFunc("/song/{ID}", controllers.FindSong).Methods(http.MethodGet)
	r.HandleFunc("/song/{ID}", controllers.DeleteSong).Methods(http.MethodDelete)
	r.HandleFunc("/test-upload", controllers.TestUpload).Methods(http.MethodPost)
	r.HandleFunc("/song/update-audio", controllers.UpdateSongAudio).Methods(http.MethodPut)
	r.HandleFunc("/song/update-cover", controllers.UpdateSongCover).Methods(http.MethodPut)
}
