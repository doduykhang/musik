package services

import (
	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(request *dto.RegisterRequest) error {
	password := []byte(request.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	result := db.Create(&models.Account{
		Email:    request.Email,
		Password: string(hashedPassword),
	})
	return result.Error
}
