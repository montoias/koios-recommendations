package users

// Movie is custom string for Movie
type Movie string

// Users is dto structure for Users
type Users struct {
	Items []User
}

// User is dto structure for User
type User struct {
	Movies []Movie
}
