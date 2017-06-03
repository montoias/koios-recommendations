package movies

// MoviesResult schema
type MoviesResult struct {
	Movies []Movie `json:"movies"`
}

// Movie schema
type Movie struct {
	Name         string   `json:"name"`
	Summary      string   `json:"summary"`
	PosterPath   string   `json:"poster_path"`
	BackdropPath string   `json:"backdrop_path"`
	ReleaseDate  string   `json:"release_date"`
	Genres       []string `json:"genres"`
	Popularity   float32  `json:"popularity"`
}
