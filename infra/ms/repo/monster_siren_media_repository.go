package repo

import (
	"fmt"
	"net/http"
)

type MonsterSirenMediaRepository struct {
	Client http.Client
}

func NewMonsterSirenMediaRepository(client http.Client) MonsterSirenMediaRepository {
	return MonsterSirenMediaRepository{client}
}

func (mediaRepository MonsterSirenMediaRepository) RetrieveAudio(url string) (response *http.Response, err error) {
	audioStream, err := mediaRepository.Client.Get(url)
	if err != nil {
		return nil, err
	}

	contentType := audioStream.Header.Get("content-type")

	switch contentType {
	case "audio/mpeg":
	case "audio/wav":
		break
	default:
		return nil, fmt.Errorf("unhandled audio/content-type retrieved: %v", contentType)
	}

	return audioStream, nil
}

func (mediaRepository MonsterSirenMediaRepository) RetrieveImage(url string) (response *http.Response, err error) {
	image, err := mediaRepository.Client.Get(url)
	if err != nil {
		return nil, err
	}

	contentType := image.Header.Get("content-type")

	switch contentType {
	case "image/jpeg":
	case "image/png":
		break
	default:
		return nil, fmt.Errorf("unhandled image/content-type retrieved: %v", contentType)
	}

	return image, nil
}

func (mediaRepository MonsterSirenMediaRepository) RetrieveLyric(url string) (response *http.Response, err error) {
	lyric, err := mediaRepository.Client.Get(url)
	if err != nil {
		return nil, err
	}

	contentType := lyric.Header.Get("content-type")

	switch contentType {
	case "application/octet-stream":
		break
	default:
		return nil, fmt.Errorf("unhandled lyric/content-type retrieved: %v", contentType)
	}

	return lyric, nil
}
