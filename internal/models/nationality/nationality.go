package nationality

type Nationality struct {
	Id   int64  `json:"-"`
	Name string `json:"name"`
}
