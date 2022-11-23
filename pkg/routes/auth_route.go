package routes

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/controllers"
	"github.com/gorilla/mux"
)


func RegisterAuthRoute(r *mux.Router){
	r.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
}
