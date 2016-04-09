package fake

func FuncMap(c *Config) map[string]interface{} {
	return map[string]interface{}{
		"rand":      func(a, b int) int { return RandInt(a, b, c) },
		"randFloat": func(a, b float64) float64 { return RandFloat(a, b, c) },
		"firstName": func() string { return FirstName(c) },
		"lastName":  func() string { return LastName(c) },
		"slice":     func(a uint64) []uint64 { return Slice(a) },
		"choice":    func(a ...interface{}) interface{} { return Choice(c, a...) },
	}
}
