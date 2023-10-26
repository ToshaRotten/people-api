package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"people-api/internal/models/person"
	"testing"
)

func TestUpdate(t *testing.T) {
	old := person.Person{
		Id: 1,
	}

	n := person.Person{
		Name:    "Adad",
		Surname: "adad",
	}

	req := Request{
		Old: old,
		New: n,
	}

	data, _ := json.Marshal(req)

	cli := http.Client{}

	r, _ := http.NewRequest(http.MethodPost, "http://localhost:8081/user/update", bytes.NewReader(data))
	resp, _ := cli.Do(r)

	respData, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respData))

}
