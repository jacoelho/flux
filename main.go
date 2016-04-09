//import "gopkg.in/Shopify/sarama.v1"

package main

import (
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"

	"github.com/jacoelho/flux/fake"
	"github.com/jacoelho/flux/payload"
)

var config = &fake.Config{
	Rand: rand.New(rand.NewSource((time.Now().Unix()))),
}

func main() {

	file, err := os.Open("test.json")
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

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
