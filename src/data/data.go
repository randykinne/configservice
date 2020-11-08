package data

import (
	"fmt"

	"github.com/randykinne/configservice/domain"
)

var dataStore map[string]*domain.Config

// Initialize data store
func Initialize() {
	dataStore = make(map[string]*domain.Config)
}

// Get a config from the map
func Get(id string) (*domain.Config, error) {
	val, ok := dataStore[id]
	if ok == false {
		return nil, fmt.Errorf("Config not found")
	}

	return val, nil
}

// Put a config in the map
func Put(config *domain.Config) (*domain.Config, error) {
	dataStore[config.ID] = config
	return config, nil
}
