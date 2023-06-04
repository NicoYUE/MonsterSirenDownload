package response

type SongResponse struct {
	SongId, Name, AlbumId, SourceUrl, LyricUrl string
	Artists                                    []string
}
