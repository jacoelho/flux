package main

import (
	"bytes"
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/jacoelho/flux/backends"
	"github.com/jacoelho/flux/fake"
	"github.com/jacoelho/flux/payload"
)

var (
	tmpPayload   bytes.Buffer
	backendNodes = flag.String("nodes", "", "backend nodes as a comma separated list")
	backendType  = flag.String("type", "kafka", "backend type: kafka|stdout")
	fileTemplate = flag.String("template", "", "template file")

	config = &fake.Config{
		Rand: rand.New(rand.NewSource((time.Now().Unix()))),
	}
)

func main() {
	flag.Parse()

	file, err := os.Open(*fileTemplate)
	if err != nil {
		log.Fatalf("open %s", err)
	}
	defer file.Close()

	p, err := payload.ReadFrom(file)
	if err != nil {
		log.Fatalf("open %s", err)
	}

	tmpl, err := template.New("").Funcs(fake.FuncMap(config)).Parse(string(p.Content()))
	if err != nil {
		log.Fatalf("template %s", err)
	}

	cfg := &backends.Config{
		Type:  *backendType,
		Nodes: strings.Split(*backendNodes, ","),
	}

	client, err := backends.NewClient(cfg)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	for {
		tmpl.Execute(&tmpPayload, nil)
		client.Serialize(tmpPayload.String())
		tmpPayload.Reset()
	}

	for {
	}
}
