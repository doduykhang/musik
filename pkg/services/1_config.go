package services

import (
	"github.com/doduykhang/musik/pkg/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = config.GetDB()
}
