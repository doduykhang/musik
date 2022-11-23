package controllers

import (
	"net/http"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/services"
	"github.com/doduykhang/musik/pkg/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var registerRequest dto.RegisterRequest
	utils.ParseBody(r, &registerRequest)
	services.Register(&registerRequest)
	utils.JsonResponse(&w, registerRequest)
}
