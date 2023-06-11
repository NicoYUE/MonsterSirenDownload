package utility

import (
	"github.com/bogem/id3v2/v2"
	"log"
	"os"
)

func PictureFrame(filename string) id3v2.PictureFrame {
	mime, _ := PictureMime(filename)
	b, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("Failed to read picture")
	}

	return id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    mime,
		PictureType: 3,
		Description: "u'Cover'",
		Picture:     b,
	}
}
