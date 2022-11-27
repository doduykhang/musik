package routes

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterAlbumRoute(r *mux.Router) {
	r.HandleFunc("/album", controllers.CreateAlbum).Methods(http.MethodPost)
	r.HandleFunc("/album", controllers.UpdateAlbum).Methods(http.MethodPut)
	r.HandleFunc("/album/update-image", controllers.UpdateAlbumImage).Methods(http.MethodPut)
	r.HandleFunc("/album/{ID}", controllers.DeleteAlbum).Methods(http.MethodDelete)
	r.HandleFunc("/album", controllers.FindAllAlbums).Methods(http.MethodGet)
	r.HandleFunc("/album/{ID}", controllers.FindAlbum).Methods(http.MethodGet)
	r.HandleFunc("/album/add-song", controllers.AddSongs).Methods(http.MethodPut)
	r.HandleFunc("/album/remove-song", controllers.RemoveSongs).Methods(http.MethodPut)
}
