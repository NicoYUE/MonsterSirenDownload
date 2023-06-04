package model

type Song struct {
	SongId, Name, AlbumName, CoverUrl, SourceUrl, LyricUrl string
	Artists                                                []string
}
