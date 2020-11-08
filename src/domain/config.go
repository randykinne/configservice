package domain

// Config structure
type Config struct {
	ID   string   `json:"id"`
	Name string   `json:"name"`
	Data []string `json:"data"`
}

// Get a config
func (c *Config) Get(name string) (map[string]interface{}, error) {
	return nil, nil
}
