package users

// UsersRequest schema
type UsersRequest struct {
	Users []User `json:"users"`
}

// User schema
type User struct {
	Movies []string `json:"movies"`
}
