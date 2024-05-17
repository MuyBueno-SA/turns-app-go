package utils

import (
	"io"
	"os"

	"github.com/pelletier/go-toml"
)

type BusinessConfig struct {
	Name          string   `toml:"name" json:"name"`
	StartTime     string   `toml:"start_time" json:"start_time"`
	EndTime       string   `toml:"end_time" json:"end_time"`
	MinModuleTime int      `toml:"min_module_time" json:"min_module_time"`
	Offices       []string `toml:"offices" json:"offices"`
}

type Config struct {
	Business BusinessConfig `toml:"business" json:"business"`
}

// LoadConfig loads the configuration from the given path.
// The configuration is expected to be in TOML format.
// The configuration is unmarshalled into a Config struct.
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = toml.Unmarshal(b, config)
	if err != nil {
		panic(err)
	}

	return config, nil
}
