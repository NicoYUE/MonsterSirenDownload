package response

type SongResponse struct {
	Cid, Name, AlbumCid, SourceUrl, LyricUrl string
	Artists                                  []string
}
