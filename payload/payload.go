package payload

import (
	"bufio"
	"bytes"
	"io"
	"text/template"
)

type payload struct {
	content string
	tmpl    *template.Template
}

type Payload interface {
	Content() string
}

func (p *payload) Content() string {
	return p.content
}

func (p *payload) parse(t map[string]interface{}) error {
	parse, err := template.New("").Funcs(t).Parse(p.content)
	if err != nil {
		return err
	}

	p.tmpl = parse
	return nil
}

func (p *payload) Generate(w io.Writer, t map[string]interface{}) error {
	if p.tmpl == nil {
		if err := p.parse(t); err != nil {
			return err
		}
	}

	if err := p.tmpl.Execute(w, nil); err != nil {
		return err
	}

	return nil
}

func (p *payload) GenerateAndLoop(w io.Writer, t map[string]interface{}) error {
	go func() {
		for {
			if err := p.Generate(w, t); err != nil {
				return
			}
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
