package services

import (
	"os"
)

type FileService interface {
	SaveFile(fileBytes []byte, path string, name string) (string, string, error)
}

type localFileService struct {
}

func GetLocalFileService() FileService {
	return &localFileService{}
}

func (service *localFileService) SaveFile(fileBytes []byte, path string, name string) (string, string, error) {
	err := os.WriteFile(path+name, fileBytes, 0644)
	if err != nil {
		return "", "", err
	}
	return path, name, nil
}
