package get_gender

import (
	"fmt"
	"testing"
)

func TestGetGender(t *testing.T) {
	gender, err := GetGender("Dmitry")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gender)
}
