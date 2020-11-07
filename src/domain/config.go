package domain

type Config struct {
	Id     int
    Name string
	config map[string]interface{}
}

// Get a config
func (c *Config) Get(name string) (map[string]interface{}, error) {
	return nil, nil
}
