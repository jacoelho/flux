package fake

import "math/rand"

type Config struct {
	Rand *rand.Rand
}

func (c *Config) getRand() *rand.Rand {
	if c.Rand == nil {
		return rand.New(rand.NewSource(42))
	}
	return c.Rand
}
