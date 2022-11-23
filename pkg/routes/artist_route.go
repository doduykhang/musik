package routes

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterArtistRoute(r *mux.Router) {
	r.HandleFunc("/artist", controllers.CreateArtist).Methods(http.MethodPost)
	r.HandleFunc("/artist", controllers.UpdateArtist).Methods(http.MethodPut)
	r.HandleFunc("/artist/{ID}", controllers.DeleteArtist).Methods(http.MethodDelete)
	r.HandleFunc("/artist", controllers.FindAllArtists).Methods(http.MethodGet)
	r.HandleFunc("/artist/{ID}", controllers.FindArtist).Methods(http.MethodGet)
}
