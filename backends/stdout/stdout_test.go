package stdout

import (
	"bytes"
	"testing"
)

func TestCreate(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Errorf("Expect nil, got %v", err)
	}
}

func TestWrite(t *testing.T) {
	source := []byte("test")
	check := new(bytes.Buffer)

	s, _ := New()

	s.setWriter(check)
	if _, err := s.Write(source); err != nil {
		t.Errorf("Expect nil, got %v", err)
	}
	if check.String() != string(source) {
		t.Errorf("Expect %v, got %v", source, check.Bytes())
	}
}
