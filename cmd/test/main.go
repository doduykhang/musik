package main

import (
	"fmt"

	"github.com/doduykhang/musik/pkg/dto"
	"github.com/doduykhang/musik/pkg/models"
	"github.com/doduykhang/musik/pkg/utils"
)

func main() {
	songDTO := dto.SongDTO{}
	song := &models.Song{}
	utils.ConverStruct(songDTO, song)
	fmt.Println(song)
}
