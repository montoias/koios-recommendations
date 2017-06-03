package movies

import "github.com/montoias/koios-recommendations/dto/movies"

// ToSchema transforms movies dto to movies result
func ToSchema(movies movies.Movies) MoviesResult {
	return MoviesResult{
		Movies: moviesToSchema(movies.Items),
	}
}

func moviesToSchema(movies []movies.Movie) []Movie {
	moviesSchema := make([]Movie, len(movies))
	for i, m := range movies {
		moviesSchema[i] = Movie{
			Name:         m.Name,
			Summary:      m.Summary,
			PosterPath:   m.PosterPath,
			BackdropPath: m.BackdropPath,
			ReleaseDate:  m.ReleaseDate.String(),
			Genres:       m.Genres,
			Popularity:   m.Popularity,
		}
	}
	return moviesSchema
}
