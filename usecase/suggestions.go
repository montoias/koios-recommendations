package usecase

import (
	"github.com/montoias/koios-recommendations/dto/movies"
	"github.com/montoias/koios-recommendations/dto/users"
	"github.com/montoias/koios-recommendations/usecase/client/tmdb"
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
	var usersMovies []movies.Movie
	for _, u := range users.Items {
		for _, m := range u.Movies {
			movie, totalResults, err := interactor.searchMovie(m)
			if err != nil {
				return movies.Movies{}, err
			}

			if totalResults == 0 {
				continue
			}

			usersMovies = append(usersMovies, movie)
		}
	}

	return movies.Movies{Items: usersMovies}, nil
}

func (interactor SuggestionsInteractor) searchMovie(movie users.Movie) (movies.Movie, int, error) {
	results, err := interactor.client.SearchMulti(string(movie), nil)
	if err != nil {
		return movies.Movie{}, results.TotalResults, err
	}

	if results.TotalResults == 0 {
		return movies.Movie{}, results.TotalResults, nil
	}

	movieResult, err := tmdb.MovieResultToMovieDto(tmdb.MovieResult(results.Results[0]))
	if err != nil {
		return movies.Movie{}, results.TotalResults, err
	}

	return movieResult, results.TotalResults, nil
}
