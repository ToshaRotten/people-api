package get_age_test

import (
	"fmt"
	"people-api/internal/http-client/handlers/get_age"
	"testing"
)

func TestGetAge(t *testing.T) {
	age, err := get_age.GetAge("Dmitry")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(age)
}
