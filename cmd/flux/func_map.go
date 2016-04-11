package main

import (
	"github.com/jacoelho/flux/fake"
	"github.com/jacoelho/flux/payload/helpers"
)

func FuncMap(c *fake.Config) map[string]interface{} {
	return map[string]interface{}{
		"rand":      func(a, b int) int { return fake.RandInt(a, b, c) },
		"randFloat": func(a, b float64) float64 { return fake.RandFloat(a, b, c) },
		"firstName": func() string { return fake.FirstName(c) },
		"lastName":  func() string { return fake.LastName(c) },
		"choice":    func(a ...interface{}) interface{} { return fake.Choice(c, a...) },
		"sequence":  func(a uint64) []uint64 { return helpers.Sequence(a) },
	}
}
