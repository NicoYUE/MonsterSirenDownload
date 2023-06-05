package main

import (
	"fmt"
	"monster-siren-record-puller/domain/model"
	"monster-siren-record-puller/domain/service"
	"monster-siren-record-puller/infra/ms/repo"
	"monster-siren-record-puller/utility"
	"net/http"
	"sync"
)

const BaseDirectory = "./MSR/"
const DefaultConcurrency = 5

var client = http.Client{}
var msrService = service.NewMonsterSirenService(client)
var mediaRepository = repo.NewMonsterSirenMediaRepository(client)

func main() {
	fmt.Printf("Hello World")

	err := utility.Mkdir(BaseDirectory, 0755)
	if err != nil {
		fmt.Println("Warn:", err)
	} else {
		fmt.Println("Created Directory: ", BaseDirectory)
	}

	//exampleUrl := "https://res01.hycdn.cn/b0f72316846ce63e1e4f6a7ad7cd965c/647D21AC/siren/audio/20230427/aaf6eb4ada3c647e8cb31b85a569c9c0.wav"
	//
	//response, err := mediaRepository.RetrieveAudio(exampleUrl)
	//utility.WriteResponse2File("DormantCraving", BaseDirectory+"DormantCraving/", response)

	albums := msrService.RetrieveAlbums()
	AsyncAlbumDownload(albums)

}

func AsyncAlbumDownload(albums []model.Album) {
	var group sync.WaitGroup
	channel := make(chan struct{}, DefaultConcurrency)

	for _, album := range albums {
		group.Add(1)
		channel <- struct{}{}

		// intermediate var
		album := album
		go func() {
			defer func() { <-channel }()
			defer group.Done()
			DownloadAlbumSongs(album)
		}()
	}
	group.Wait()
}

func DownloadAlbumSongs(album model.Album) {
	fmt.Printf("Album: %s\n", album.Name)
	songs := msrService.RetrieveAlbumSongs(album)
	albumPath := fmt.Sprintf(BaseDirectory+"%s/", album.Name)
	utility.Mkdir(albumPath, 0755)

	coverResp, _ := mediaRepository.RetrieveImage(album.CoverUrl)
	utility.WriteResponse2File(album.Name, albumPath, coverResp)

	for _, song := range songs {
		fmt.Printf("Downloading : %s\n", song.Name)

		audioResp, _ := mediaRepository.RetrieveAudio(song.SourceUrl)
		utility.WriteResponse2File(song.Name, albumPath, audioResp)

		if song.LyricUrl != "" {
			lyricResp, _ := mediaRepository.RetrieveLyric(song.LyricUrl)
			utility.WriteResponse2File(song.Name, albumPath, lyricResp)
		}

		fmt.Printf("---- Finished : %s\n", song.Name)
	}
}
