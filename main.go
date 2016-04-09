//import "gopkg.in/Shopify/sarama.v1"

package main

import (
	"flux/generator"
	"flux/payload"
	"fmt"
	"log"
	"math/rand"
	"os"
	"text/template"
	"time"
)

var config = &generator.Config{
	Rand: rand.New(rand.NewSource((time.Now().Unix()))),
}

func main() {

	funcMap := template.FuncMap{
		"rand":      func(a, b int) int { return generator.RandInt(a, b, config) },
		"randFloat": func(a, b float64) float64 { return generator.RandFloat(a, b, config) },
	}
	templateText := `
    Output 1: "{{ randFloat 10.01 15.02 | printf "%.2f" }}"
`
	tmpl := template.Must(template.New("titleTest").Funcs(funcMap).Parse(templateText))

	// Run the template to verify the output.
	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

	file, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("open %s", err)
	}
	defer file.Close()

	p, err := payload.ReadFrom(file)
	if err != nil {
		log.Fatalf("open %s", err)
	}
	fmt.Println(string(p.Content()))
}
