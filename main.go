package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jacoelho/flux/backends"
	"github.com/jacoelho/flux/fake"
	"github.com/jacoelho/flux/payload"
)

var (
	backendNodes = flag.String("nodes", "", "backend nodes as a comma separated list")
	backendType  = flag.String("type", "stdout", "backend type: kafka|stdout")
	fileTemplate = flag.String("template", "", "template file")
	duration     = flag.Int("duration", 10, "number of seconds running")
)

func main() {
	flag.Parse()

	cfg := &backends.Config{
		Type:  *backendType,
		Nodes: strings.Split(*backendNodes, ","),
	}

	config := &fake.Config{
		Rand: rand.New(rand.NewSource((time.Now().Unix()))),
	}

	client, err := backends.NewClient(cfg)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	file, err := os.Open(*fileTemplate)
	if err != nil {
		log.Fatalf("open %s", err)
	}
	defer file.Close()

	p, err := payload.ReadFrom(file)
	if err != nil {
		log.Fatalf("open %s", err)
	}

	timeout := make(chan struct{}, 1)
	go func() {
		time.Sleep(time.Second * time.Duration(*duration))
		timeout <- struct{}{}
	}()

	p.GenerateAndLoop(client, FuncMap(config))

	<-timeout
}
