package string

import "testing"

func TestGetSides(t *testing.T) {
	value := "ThisIsStringThat"

	result := GetSides(value, 3)
	t.Log(result)
}
