package stdout

import (
	"io"
	"os"
)

type stdout struct {
	writer io.Writer
}

func New() (*stdout, error) {
	return &stdout{writer: os.Stdout}, nil
}

func (s *stdout) setWriter(w io.Writer) {
	s.writer = w
}

func (s *stdout) Write(p []byte) (n int, err error) {
	return s.writer.Write(p)
}

func (s *stdout) Flush() error {
	return nil
}

func (s *stdout) Close() error {
	return nil
}
