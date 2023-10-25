package pagination

type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	From       int `json:"from"`
	To         int `json:"to"`
	MaxPerPage int `json:"max_per_page"`
	Total      int `json:"total"`
}

func Calculate(page int, perPage int, total int) (int, int) {
	offset := (page - 1) * perPage
	return perPage, offset
}
