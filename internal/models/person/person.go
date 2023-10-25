package person

type Person struct {
	Id          int64  `json:"-"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int8   `json:"age"`
	Sex         string `json:"sex"`
	Nationality string `json:"nationality"`
}
