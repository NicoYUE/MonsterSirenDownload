package service

import (
	"monster-siren-record-puller/domain/model"
	"monster-siren-record-puller/infra/ms/repo"
	"net/http"
	"strings"
)

type MonsterSirenService struct {
	Client                 http.Client
	MonsterSirenRepository repo.MonsterSirenRepository
}

func NewMonsterSirenService(client http.Client) MonsterSirenService {
	return MonsterSirenService{
		Client:                 client,
		MonsterSirenRepository: repo.NewMonsterSirenRepository(client),
	}
}

func (service MonsterSirenService) RetrieveAlbums() []model.Album {
	albumResponse := service.MonsterSirenRepository.RetrieveAlbums()

	albums := make([]model.Album, len(albumResponse))
	for idx, obj := range albumResponse {
		trimmedName := strings.TrimSpace(obj.Name)
		album := model.Album{
			AlbumId:  obj.Cid,
			Name:     trimmedName,
			CoverUrl: obj.CoverUrl,
		}
		albums[idx] = album
	}

	return albums
}

func (service MonsterSirenService) RetrieveSong(album model.Album, songId string) model.Song {
	songResponse := service.MonsterSirenRepository.RetrieveSong(songId)

	return model.Song{
		SongId:    songResponse.SongId,
		Name:      songResponse.Name,
		AlbumName: album.Name,
		CoverUrl:  album.CoverUrl,
		SourceUrl: songResponse.SourceUrl,
		LyricUrl:  songResponse.LyricUrl,
		Artists:   songResponse.Artists,
	}
}

func (service MonsterSirenService) RetrieveAlbumSongs(album model.Album) []model.Song {
	albumDetailsResponse := service.MonsterSirenRepository.RetrieveAlbumDetails(album.AlbumId)

	songs := make([]model.Song, len(albumDetailsResponse.Songs))
	for idx, songId := range albumDetailsResponse.Songs {
		songs[idx] = service.RetrieveSong(album, songId.Cid)
	}

	return songs
}
