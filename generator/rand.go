package generator

func RandInt(min int, max int, c *Config) int {
	return c.getRand().Intn(max-min) + min
}

func RandFloat(min float64, max float64, c *Config) float64 {
	return c.getRand().Float64()*(max-min) + min
}
