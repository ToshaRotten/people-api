package save

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"people-api/internal/models/person"
	"testing"
)

func TestSave(t *testing.T) {

	p := person.Person{
		Name:       "Ad",
		Surname:    "Ad",
		Patronymic: "Ad",
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}

	cli := http.Client{}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8081/user/save", bytes.NewReader(data))

	resp, err := cli.Do(req)

	respData, err := io.ReadAll(resp.Body)

	fmt.Println(string(respData))
}
