package users

// Users schema
type Users struct {
	Items []User `json:"users"`
}

// User schema
type User struct {
	Movies []string `json:"movies"`
}
