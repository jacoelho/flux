package fake

func Choice(c *Config, values ...interface{}) interface{} {
	if len(values) != 0 {
		return values[c.getRand().Intn(len(values))]
	}
	return nil
}
