package stdout

import "os"

type stdout struct{}

func New() (*stdout, error) {
	return &stdout{}, nil
}

func (s *stdout) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (s *stdout) Flush() error {
	return nil
}

func (s *stdout) Close() error {
	return nil
}
