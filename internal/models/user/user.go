package user

type User struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         string `json:"age"`
	Sex         string `json:"sex"`
	Nationality string `json:"nationality"`
}
