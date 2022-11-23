package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, x interface{}) error {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, x)
	return err
}

func JsonResponse(w *http.ResponseWriter, x interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(x)
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(res)
}

func ConverseStruct[S any, D any](source S, destination D) {
	jsonRes, _ := json.Marshal(source)
	json.Unmarshal(jsonRes, destination)
}

func GetIDFromRequest(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	IDRaw := vars["ID"]
	ID64, err := strconv.ParseUint(IDRaw, 10, 32)
	if err != nil {
		return 0, err
	}
	ID := uint(ID64)
	return ID, err
}
