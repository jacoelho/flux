package backends

type Config struct {
	Type  string
	Nodes []string
	Topic string // leaking kafka configuration
}
