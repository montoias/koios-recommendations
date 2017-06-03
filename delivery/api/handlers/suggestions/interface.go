package suggestions

import (
	"github.com/montoias/koios-recommendations/dto/movies"
	"github.com/montoias/koios-recommendations/dto/users"
)

// SuggestionsInteractor interface
type SuggestionsInteractor interface {
	CreateSuggestions(users.Users) (movies.Movies, error)
}
