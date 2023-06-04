package response

type AlbumDetailsResponse struct {
	Cid   string `json:"cid"`
	Songs []struct {
		Cid string `json:"cid"`
	} `json:"songs"`
}
