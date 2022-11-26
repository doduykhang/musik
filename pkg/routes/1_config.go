package routes

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	RegisterAuthRoute(r)
	RegisterArtistRoute(r)
	RegisterSongRoute(r)
}
