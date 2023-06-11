package model

import "github.com/bogem/id3v2/v2"

type SongMetadata struct {
	Title, AlbumName      string
	AlbumArtists, Artists []string
	PictureFrame          id3v2.PictureFrame
}
