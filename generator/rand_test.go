package generator

import "testing"

func TestRandInt(t *testing.T) {
	c := &Config{}

	if RandInt(10, 15, c) != 10 {
		t.Error("Wrong value")
	}
}

func TestRandFloat(t *testing.T) {
	c := &Config{}

	if RandFloat(10, 15, c) != 11.865141805233163 {
		t.Error("Wrong value")
	}
}
