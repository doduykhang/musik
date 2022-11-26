package controllers

import (
	mapper "github.com/dranikpg/dto-mapper"
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

func Map(to, from interface{}) error {
	return mapper.Map(to, from)
}
