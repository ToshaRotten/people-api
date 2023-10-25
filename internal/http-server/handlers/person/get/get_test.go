package get

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	cli := http.Client{}

	r := Request{Id: 2}

	data, err := json.Marshal(r)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/user/get", bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.Do(req)

	respData, err := io.ReadAll(resp.Body)

	fmt.Println(string(respData))
}
