package generator

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
