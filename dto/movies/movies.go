package movies

import "time"

// Movies is dto structure for Movies
type Movies struct {
	Items []Movie
}

// Movie is dto structure for Movie
type Movie struct {
	Name         string
	Summary      string
	PosterPath   string
	BackdropPath string
	ReleaseDate  time.Time
	Genres       []string
	Popularity   float32
}
