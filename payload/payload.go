package payload

import (
	"bufio"
	"bytes"
	"io"
)

type payload struct {
	content []byte
}

type Payload interface {
	Content() []byte
}

func (p *payload) Content() []byte {
	return p.content
}

func extractContent(r io.Reader) (content []byte, err error) {
	wr := new(bytes.Buffer)
	if _, err = wr.ReadFrom(r); err != nil {
		return
	}
	return wr.Bytes(), nil
}

func ReadFrom(r io.Reader) (p *payload, err error) {
	reader := bufio.NewReader(r)
	newPayload := new(payload)

	newPayload.content, err = extractContent(reader)

	if err != nil {
		return nil, err
	}

	return newPayload, nil
}
