package routes

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterPlaylistRoute(r *mux.Router) {
	r.HandleFunc("/playlist", controllers.CreatePlaylist).Methods(http.MethodPost)
	r.HandleFunc("/playlist", controllers.UpdatePlaylist).Methods(http.MethodPut)
	r.HandleFunc("/playlist/user/{ID}", controllers.FindPlaylists).Methods(http.MethodGet)
	r.HandleFunc("/playlist/{ID}", controllers.DeletePlaylist).Methods(http.MethodDelete)
	r.HandleFunc("/playlist/{ID}", controllers.FindPlaylist).Methods(http.MethodGet)
	r.HandleFunc("/playlist/add-song", controllers.AddSongToPlaylist).Methods(http.MethodPut)
	r.HandleFunc("/playlist/remove-song", controllers.RemoveSongFromPlaylist).Methods(http.MethodPut)
}
