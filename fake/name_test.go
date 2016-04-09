package fake

import "testing"

func TestFirstName(t *testing.T) {
	c := &Config{}

	if FirstName(c) != "Mia" {
		t.Error("Wrong value")
	}
}

func TestLastName(t *testing.T) {
	c := &Config{}

	if LastName(c) != "Davies" {
		t.Error("Wrong value")
	}
}
