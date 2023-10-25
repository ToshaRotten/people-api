package save_many

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"people-api/internal/models/person"
	"testing"
)

func TestSaveMany(t *testing.T) {

	p := []person.Person{
		{
			Name:       "Dmitriy",
			Surname:    "Ushakov",
			Patronymic: "Vasilevich",
		},
		{
			Name:       "qeq",
			Surname:    "qqeq",
			Patronymic: "qqqq",
		},
	}

	data, err := json.Marshal(Request{Items: p})
	if err != nil {
		t.Error(err)
	}

	cli := http.Client{}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8081/users/save", bytes.NewReader(data))

	resp, err := cli.Do(req)

	respData, err := io.ReadAll(resp.Body)

	fmt.Println(string(respData))
}
