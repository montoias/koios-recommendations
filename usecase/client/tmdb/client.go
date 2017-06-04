package tmdb

import (
	"fmt"
	"time"

	"github.com/montoias/koios-recommendations/dto/movies"
	gotmdb "github.com/ryanbradynd05/go-tmdb"
)

const (
	timeFormatLayout = "2006-01-02"
	baseImageURL     = "https://image.tmdb.org/t/p/w500"
)

// Client is interface for the themoviedb.org API client
type Client interface {
	SearchMulti(name string, options map[string]string) (*gotmdb.MultiSearchResults, error)
}

// MovieResultToMovieDto transforms gotmdb MovieShort to Movie Dto
func MovieResultToMovieDto(movieResult MovieResult) (movies.Movie, error) {
	var releaseDate time.Time
	var err error

	if movieResult.FirstAirDate != "" {
		releaseDate, err = time.Parse(timeFormatLayout, movieResult.FirstAirDate)
		if err != nil {
			return movies.Movie{}, err
		}
	}

	return movies.Movie{
		Name:         movieResult.OriginalTitle,
		Summary:      movieResult.Overview,
		Popularity:   movieResult.Popularity,
		ReleaseDate:  releaseDate,
		PosterPath:   buildImagePath(movieResult.PosterPath),
		BackdropPath: buildImagePath(movieResult.BackdropPath),
	}, nil
}

func buildImagePath(uri string) string {
	return fmt.Sprintf("%s%s", baseImageURL, uri)
}

type MovieResult struct {
	BackdropPath  string `json:"backdrop_path"`
	ID            int
	OriginalName  string   `json:"original_name"`
	OriginalTitle string   `json:"original_title"`
	Overview      string   `json:"overview"`
	FirstAirDate  string   `json:"first_air_date"`
	OriginCountry []string `json:"origin_country"`
	PosterPath    string   `json:"poster_path"`
	Popularity    float32
	Name          string
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     uint32  `json:"vote_count"`
	MediaType     string  `json:"media_type"`
}
