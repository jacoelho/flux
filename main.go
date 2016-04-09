//import "gopkg.in/Shopify/sarama.v1"

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	funcMap := template.FuncMap{
		"title": strings.Title,
	}
	templateText := `
    Output 0: {{title "adsa" }}
    Output 1: {{title "asda asdad" "123131"}}
`
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}
	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
