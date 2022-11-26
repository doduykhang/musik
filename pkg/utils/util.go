package utils

import (
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/doduykhang/musik/pkg/dto"
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

func ErrorResponse(w *http.ResponseWriter, message string, code int) {
	(*w).Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(
		dto.ErrorResponse{
			Code:    code,
			Message: message,
		},
	)
	(*w).WriteHeader(code)
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
	return ID, nil
}

func GetPagination(r *http.Request) (*dto.Pagination, error) {
	q := r.URL.Query()
	page, err := strconv.Atoi(q.Get("page"))

	if err != nil {
		return nil, err
	}

	if page == 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(q.Get("page_size"))
	if err != nil {
		return nil, err
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	return &dto.Pagination{
		Page: page,
		Size: pageSize,
	}, nil
}
func GetFile(r *http.Request, field string) ([]byte, *multipart.FileHeader, error) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile(field)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	return fileBytes, handler, nil
}

func GetFileByteWithName(r *http.Request, field string) ([]byte, string, error) {
	bytes, handler, err := GetFile(r, field)
	if err != nil {
		return nil, "", err
	}
	return bytes, handler.Filename, err
}
