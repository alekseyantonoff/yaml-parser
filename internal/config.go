package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents a parsed YAML configuration as a key-value map.
// The keys are strings, and values can be any YAML type (map, list, string, number, etc.).
type Config map[string]interface{}

// LoadConfig reads a YAML file from the given filePath and parses it into a Config object.
// It returns the parsed Config and any error encountered during reading or unmarshaling.
func LoadConfig(filePath string) (Config, error) {
	// Read the entire file into memory
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal (parse) the YAML data into the Config map
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
