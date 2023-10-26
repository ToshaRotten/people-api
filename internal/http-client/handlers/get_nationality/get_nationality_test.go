package get_nationality

import (
	"fmt"
	"testing"
)

func TestGetNationality(t *testing.T) {
	nationality, err := GetNationality("Dmitry")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(nationality)
}
