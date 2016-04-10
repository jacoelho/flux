package payload

import (
	"bytes"
	"testing"
)

func TestLoadContent(t *testing.T) {
	value := bytes.NewBufferString("test")

	p, err := ReadFrom(value)
	if err != nil {
		t.Errorf("Expect nil, got %v", err)
	}

	if p.Content() != "test" {
		t.Errorf("Expect test, got %v", p.Content())
	}
}

func TestParseValid(t *testing.T) {
	value := bytes.NewBufferString("{{ sum 2 3 }}")

	p, _ := ReadFrom(value)

	result := new(bytes.Buffer)
	f := map[string]interface{}{
		"sum": func(a, b int) int { return a + b },
	}

	err := p.Generate(result, f)
	if err != nil {
		t.Errorf("Expect nil, got %v", err)
	}

	if result.String() != "5" {
		t.Errorf("Expected 5, got %s", result.String())
	}
}

func TestParseInvalid(t *testing.T) {
	value := bytes.NewBufferString("{{ sum 2 3 }}")

	p, _ := ReadFrom(value)

	result := new(bytes.Buffer)
	f := map[string]interface{}{}

	err := p.Generate(result, f)
	if err == nil {
		t.Errorf("Expect error, got %v", err)
	}
}

func TestParseEmpty(t *testing.T) {
	value := bytes.NewBufferString("")

	p, _ := ReadFrom(value)

	result := new(bytes.Buffer)
	f := map[string]interface{}{}

	err := p.Generate(result, f)
	if err != nil {
		t.Errorf("Expect nil, got %v", err)
	}

	if result.String() != "" {
		t.Errorf("Expected \"\", got %s", result.String())
	}
}
