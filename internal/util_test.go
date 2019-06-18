package internal

import "testing"

func TestPrettify(t *testing.T) {
	type Person struct {
		Age    int
		Name   *string
		School map[string]*string
	}

	school := make(map[string]*string)
	school["junior"] = String("junior school")
	school["high"] = String("high school")
	person := Person{
		Age:    1,
		Name:   String("jess"),
		School: school,
	}

	t.Log(Prettify(person))
}
