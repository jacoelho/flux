package fake

import "testing"

func TestRandInt(t *testing.T) {
	c := &Config{}

	if val := RandInt(10, 15, c); val != 10 {
		t.Errorf("Wrong value, got %v", val)
	}
}

func TestRandFloat(t *testing.T) {
	c := &Config{}

	if val := RandFloat(10, 15, c); val != 11.865141805233163 {
		t.Errorf("Wrong value, got %v", val)
	}
}
