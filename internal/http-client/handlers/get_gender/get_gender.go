package get_gender

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GenderizeResponse struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
	Count       int     `json:"count"`
}

func GetGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var genderizeResponse GenderizeResponse
	err = json.NewDecoder(resp.Body).Decode(&genderizeResponse)
	if err != nil {
		return "", err
	}

	if genderizeResponse.Gender == "male" {
		return "M", err
	}
	if genderizeResponse.Gender == "female" {
		return "F", err
	}

	return "", nil
}
