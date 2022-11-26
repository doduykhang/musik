package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

var (
	validate *validator.Validate
	decoder  *schema.Decoder
)

func init() {
	validate = validator.New()
	decoder = schema.NewDecoder()
}
