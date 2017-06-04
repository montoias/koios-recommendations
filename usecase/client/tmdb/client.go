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
	SearchMovie(name string, options map[string]string) (*gotmdb.MovieSearchResults, error)
}

// MovieShortToMovieDto transforms gotmdb MovieShort to Movie Dto
func MovieShortToMovieDto(short gotmdb.MovieShort) (movies.Movie, error) {
	releaseDate, err := time.Parse(timeFormatLayout, short.ReleaseDate)
	if err != nil {
		return movies.Movie{}, err
	}

	return movies.Movie{
		Name:         short.Title,
		Popularity:   short.Popularity,
		ReleaseDate:  releaseDate,
		PosterPath:   buildImagePath(short.PosterPath),
		BackdropPath: buildImagePath(short.BackdropPath),
	}, nil
}

func buildImagePath(uri string) string {
	return fmt.Sprintf("%s%s", baseImageURL, uri)
}
