package generator

import (
	"math/rand"
	"time"
)

type Config struct {
	Rand *rand.Rand
}

func (c *Config) getRand() *rand.Rand {
	if c.Rand == nil {
		seed := rand.NewSource(time.Now().UnixNano())
		return rand.New(seed)
	}
	return c.Rand
}

func FirstName(c *Config) string {
	first := [...]string{
		"Sophia",
		"Emma",
		"Olivia",
		"Ava",
		"Isabella",
		"Mia",
		"Zoe",
		"Lily",
		"Emily",
		"Madelyn",
		"Jackson",
		"Aiden",
		"Liam",
		"Lucas",
		"Noah",
		"Mason",
		"Ethan",
		"Caden",
		"Jacob",
		"Logan",
	}
	return first[c.getRand().Intn(len(first))]
}

func LastName(c *Config) string {
	last := [...]string{
		"Smith",
		"Jones",
		"Williams",
		"Taylor",
		"Brown",
		"Davies",
		"Evans",
		"Wilson",
		"Thomas",
		"Johnson",
		"Roberts",
		"Robinson",
		"Thompson",
		"Wright",
		"Walker",
		"White",
		"Edwards",
		"Hughes",
		"Green",
		"Hall",
	}
	return last[c.getRand().Intn(len(last))]
}

func randInt(min int, max int, c *Config) int {
	return c.getRand().Intn(max-min) + min
}

func randFloat(min float64, max float64, c *Config) float64 {
	return c.getRand().Float64()*(max-min) + min
}
