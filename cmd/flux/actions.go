package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jacoelho/flux/backends"
	"github.com/jacoelho/flux/fake"
	"github.com/jacoelho/flux/payload"
)

type ActionConfig struct {
	Type             string
	ConnectionString string
	Topic            string
	Template         string
	Duration         time.Duration
}

func execute(a *ActionConfig) {
	cfg := &backends.Config{
		Type:  a.Type,
		Nodes: strings.Split(a.ConnectionString, ","),
		Topic: a.Topic,
	}

	config := &fake.Config{
		Rand: rand.New(rand.NewSource((time.Now().Unix()))),
	}

	client, err := backends.NewClient(cfg)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	file, err := os.Open(a.Template)
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
		time.Sleep(time.Second * a.Duration)
		timeout <- struct{}{}
	}()

	p.GenerateAndLoop(client, FuncMap(config))

	<-timeout

}
