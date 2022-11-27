package services

func GetAlbumService() AlbumService {
	return &albumServiceImpl{}
}

func GetSongService() SongService {
	return &songServiceImpl{}
}

func GetPlaylistService() PlaylistService {
	return &playlistServiceImpl{}
}
