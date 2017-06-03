package users

import usersDto "github.com/montoias/koios-recommendations/dto/users"

// ToDto transforms users schema to dto
func ToDto(users Users) usersDto.Users {
	ui := make([]usersDto.User, len(users.Items))
	for i, s := range users.Items {
		ui[i] = usersDto.User{Movies: moviesToDto(s.Movies)}
	}
	return usersDto.Users{
		Items: ui,
	}
}

func moviesToDto(movies []string) []usersDto.Movie {
	m := make([]usersDto.Movie, len(movies))
	for i, s := range movies {
		m[i] = usersDto.Movie(s)
	}
	return m
}
