package usecase

import (
	"github.com/montoias/koios-recommendations/dto/movies"
	"github.com/montoias/koios-recommendations/dto/users"
	"github.com/montoias/koios-recommendations/usecase/client/tmdb"
	gotmdb "github.com/ryanbradynd05/go-tmdb"
)

// SuggestionsInteractor provides methods for suggestions
type SuggestionsInteractor struct {
	client tmdb.Client
}

// NewSuggestionsInteractor returns initialized pointer to SuggestionsInteractor
func NewSuggestionsInteractor(client tmdb.Client) *SuggestionsInteractor {
	return &SuggestionsInteractor{client: client}
}

// CreateSuggestions creates a list of movie suggestions based on a list of users
func (interactor SuggestionsInteractor) CreateSuggestions(users users.Users) (movies.Movies, error) {
	var zeroValueMovieShort gotmdb.MovieShort
	var usersMovies []movies.Movie
	for _, u := range users.Items {
		for _, m := range u.Movies {
			movieShort, err := interactor.searchMovie(m)
			if err != nil {
				return movies.Movies{}, err
			}

			if movieShort == zeroValueMovieShort {
				continue
			}

			movieDto, err := tmdb.MovieShortToMovieDto(movieShort)
			if err != nil {
				return movies.Movies{}, err
			}
			usersMovies = append(usersMovies, movieDto)
		}
	}

	return movies.Movies{Items: usersMovies}, nil
}

func (interactor SuggestionsInteractor) searchMovie(movie users.Movie) (gotmdb.MovieShort, error) {
	results, err := interactor.client.SearchMovie(string(movie), nil)
	if err != nil {
		return gotmdb.MovieShort{}, err
	}

	if results.TotalResults == 0 {
		return gotmdb.MovieShort{}, nil
	}

	return results.Results[0], nil
}
