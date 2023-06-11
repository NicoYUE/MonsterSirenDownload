package utility

import (
	"fmt"
	"github.com/NicoYUE/godub/converter"
	"github.com/bogem/id3v2/v2"
	"log"
	"monster-siren-record-puller/domain/model"
	"os"
	"path/filepath"
	"strings"
)

func SetID3Tags(filename string, metadata model.SongMetadata) {
	tag, _ := id3v2.Open(filename, id3v2.Options{Parse: true})
	defer tag.Close()

	tag.SetTitle(metadata.Title)
	tag.SetAlbum(metadata.AlbumName)
	tag.SetArtist(strings.Join(metadata.Artists, ","))
	tag.AddAttachedPicture(metadata.PictureFrame)

	err := tag.Save()
	if err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}
}

func PictureMime(filename string) (string, error) {
	extension := filepath.Ext(filename)

	switch extension {
	case ".jpg":
		return "image/jpeg", nil
	case ".png":
		return "image/png", nil
	}
	return "", fmt.Errorf("currently unhandled image type %s", extension)
}

func WavToMp3(filename string) string {
	mp3Filename := filename[0:len(filename)-4] + ".mp3"

	wavFile, _ := os.Open(filename)
	// Create if not exist, write only, truncate content
	mp3File, _ := os.OpenFile(mp3Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)

	err := converter.NewConverter(mp3File).WithBitRate(128000).WithDstFormat("mp3").Convert(wavFile)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(filename)
	defer wavFile.Close()
	defer mp3File.Close()

	return mp3Filename
}
