package payload

import (
	"bufio"
	"bytes"
	"html/template"
	"io"
)

type payload struct {
	content         string
	transformations map[string]interface{}
}

type Payload interface {
	Content() string
}

func (p *payload) Content() string {
	return p.content
}

func (p *payload) Generate(w io.Writer, t map[string]interface{}) error {
	tmpl, err := template.New("").Funcs(t).Parse(p.content)
	if err != nil {
		return err
	}

	go func() {
		for {
			tmpl.Execute(w, nil)
		}
	}()

	return nil
}

func extractContent(r io.Reader) (content string, err error) {
	wr := new(bytes.Buffer)
	if _, err = wr.ReadFrom(r); err != nil {
		return
	}
	return wr.String(), nil
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
