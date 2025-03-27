package string

import "testing"

func TestGetSides(t *testing.T) {
	value := "ThisIsStringThat"

	result := GetSides(value, 3)
	t.Log(result)
}

func TestNominateFileName(t *testing.T) {
	value := "ThisIsStringThat"
	result := NominateFileName(value)
	t.Log(result)

	value = "ThisIsStringThat.xlsx"
	result = NominateFileName(value)
	t.Log(result)

	value = "ThisIsStringThat (1).xlsx"
	result = NominateFileName(value)
	t.Log(result)

	value = "ThisIsStringThat (2).xlsx"
	result = NominateFileName(value)
	t.Log(result)

	value = ".gitignore"
	result = NominateFileName(value)
	t.Log(result)

	value = "1.xlsx"
	result = NominateFileName(value)
	t.Log(result)

	value = "1..xlsx"
	result = NominateFileName(value)
	t.Log(result)
}
