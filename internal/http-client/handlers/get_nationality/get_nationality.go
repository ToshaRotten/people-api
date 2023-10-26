package get_nationality

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NationalizeResponse struct {
	Name    string `json:"name"`
	Country []struct {
		CountryId   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func GetNationality(name string) ([]string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var nationalizeResponse NationalizeResponse
	err = json.NewDecoder(resp.Body).Decode(&nationalizeResponse)
	if err != nil {
		return nil, err
	}

	countries := make([]string, len(nationalizeResponse.Country))
	for i, c := range nationalizeResponse.Country {
		countries[i] = c.CountryId
	}

	return countries, nil
}
