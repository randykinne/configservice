package domain

type Config struct {
	config map[string]interface{}
}


func (c *Config) Get(name string) (map[string]interface{}, error) {
	
}