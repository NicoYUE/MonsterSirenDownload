package repo

import (
	"encoding/json"
	"fmt"
	"io"
	"monster-siren-record-puller/infra/ms/response"
	"net/http"
)

const (
	AlbumsUrl       = "https://monster-siren.hypergryph.com/api/albums"
	SongUrl         = "https://monster-siren.hypergryph.com/api/song/%s"
	AlbumDetailsUrl = "https://monster-siren.hypergryph.com/api/album/%s/detail"
)

type MonsterSirenRepository struct {
	Client http.Client
}

func NewMonsterSirenRepository(client http.Client) MonsterSirenRepository {
	return MonsterSirenRepository{client}
}

func (repository MonsterSirenRepository) RetrieveAlbums() []response.AlbumResponse {
	jsonData := repository.RetrieveRawJsonMsrData(AlbumsUrl)

	var albumResponse []response.AlbumResponse
	err := json.Unmarshal(jsonData, &albumResponse)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return albumResponse
}

func (repository MonsterSirenRepository) RetrieveAlbumDetails(albumId string) response.AlbumDetailsResponse {
	url := fmt.Sprintf(AlbumDetailsUrl, albumId)
	jsonData := repository.RetrieveRawJsonMsrData(url)

	var detailsResponse response.AlbumDetailsResponse
	err := json.Unmarshal(jsonData, &detailsResponse)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return detailsResponse
}

func (repository MonsterSirenRepository) RetrieveSong(songId string) response.SongResponse {
	url := fmt.Sprintf(SongUrl, songId)
	jsonData := repository.RetrieveRawJsonMsrData(url)

	var songResponse response.SongResponse
	err := json.Unmarshal(jsonData, &songResponse)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return songResponse
}

func (repository MonsterSirenRepository) RetrieveRawJsonMsrData(url string) json.RawMessage {
	data, err := repository.Client.Get(url)
	if err != nil {
		fmt.Printf(err.Error())
	}

	responseData, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Printf(err.Error())
	}

	var msrResponse response.MsrResponse
	err = json.Unmarshal(responseData, &msrResponse)
	if err != nil {
		return nil
	}

	return msrResponse.Data
}
