package services

import (
	"github.com/doduykhang/musik/pkg/config"
	mapper "github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = config.GetDB()
}

func Map(to, from interface{}) error {
	return mapper.Map(to, from)
}
