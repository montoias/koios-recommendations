package tmdb

import gotmdb "github.com/ryanbradynd05/go-tmdb"

// Client is themoviedb.org API wrapper
type Client struct {
	*gotmdb.TMDb
}

// New returns initialized instance of Client
func New(apiKey string) *Client {
	t := gotmdb.Init(apiKey)
	return &Client{t}
}

func (t *Client) Test() *gotmdb.Movie {
	fightClubInfo, _ := t.GetMovieInfo(550, nil)
	return fightClubInfo
}
