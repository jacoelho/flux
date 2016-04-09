package stdout

import "fmt"

type stdout struct{}

func New() (*stdout, error) {
	return &stdout{}, nil
}

func (s *stdout) Serialize(msg string) error {
	fmt.Printf("message: %s\n", msg)
	return nil
}

func (s *stdout) Close() error {
	return nil
}
