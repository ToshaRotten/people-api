package person

import "people-api/internal/models/nationality"

type Person struct {
	Id          int64                   `json:"-"`
	Name        string                  `json:"name"`
	Surname     string                  `json:"surname"`
	Patronymic  string                  `json:"patronymic"`
	Age         string                  `json:"age"`
	Sex         string                  `json:"sex"`
	Nationality nationality.Nationality `json:"nationality"`
}
