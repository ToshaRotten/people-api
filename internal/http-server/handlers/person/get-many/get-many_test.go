package get_many

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestGetMany(t *testing.T) {
	cli := http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/user/get?page=1&per_page=100", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.Do(req)

	respData, err := io.ReadAll(resp.Body)

	fmt.Println(string(respData))
}
