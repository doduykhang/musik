package services

func GetAlbumService() AlbumService {
	return &albumServiceImpl{}
}
