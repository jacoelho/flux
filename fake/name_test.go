package fake

import "testing"

func TestFirstName(t *testing.T) {
	c := &Config{}

	if val := FirstName(c); val != "Mia" {
		t.Errorf("Wrong value, got %v", val)
	}
}

func TestLastName(t *testing.T) {
	c := &Config{}

	if val := LastName(c); val != "Davies" {
		t.Errorf("Wrong value, got %v", val)
	}
}
