package helpers

import "testing"

func TestSequence(t *testing.T) {

	if check := Sequence(10); len(check) != 10 {
		t.Errorf("wrong value, got %v", check)
	}
}
