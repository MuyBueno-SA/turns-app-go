package utils

import (
	"io"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Business struct {
		Name          string   `toml:"name" json:"name"`
		StartTime     string   `toml:"start_time" json:"start_time"`
		EndTime       string   `toml:"end_time" json:"end_time"`
		MinModuleTime int      `toml:"min_module_time" json:"min_module_time"`
		Offices       []string `toml:"offices" json:"offices"`
	} `toml:"business" json:"business"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	config := &Config{}
	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(b, config)
	if err != nil {
		panic(err)
	}

	return config, nil
}
