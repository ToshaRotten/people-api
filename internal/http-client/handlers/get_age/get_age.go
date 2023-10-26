package get_age

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AgifyResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

func GetAge(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var agifyResponse AgifyResponse
	err = json.NewDecoder(resp.Body).Decode(&agifyResponse)
	if err != nil {
		return 0, err
	}

	return agifyResponse.Age, nil
}
