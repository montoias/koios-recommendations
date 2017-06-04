package users

// UsersRequest schema
type UsersRequest struct {
	Users []User `json:"users"`
}

// User schema
type User struct {
	ID     string   `json:"id"`
	Movies []string `json:"movies"`
}
