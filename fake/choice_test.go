package fake

import "testing"

func TestChoiceString(t *testing.T) {
	c := &Config{}

	if check := Choice(c, "one", "two"); check != "two" {
		t.Errorf("wrong value, got %v", check)
	}
}

func TestChoiceInt(t *testing.T) {
	c := &Config{}

	if check := Choice(c, 1, 2); check != 2 {
		t.Errorf("wrong value, got %v", check)
	}
}
