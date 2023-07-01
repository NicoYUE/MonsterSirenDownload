package utility

import (
	"github.com/bogem/id3v2/v2"
	"io"
	"log"
	"os"
)

func PictureFrame(filename string) id3v2.PictureFrame {
	mime, _ := PictureMime(filename)
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("Failed to read picture")
	}

	return id3v2.PictureFrame{
		Encoding:    id3v2.EncodingISO,
		MimeType:    mime,
		PictureType: id3v2.PTFrontCover,
		Description: "Cover",
		Picture:     data,
	}
}
