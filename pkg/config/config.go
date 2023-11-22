package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

var configuration Config

func Get() Config {
	return configuration
}

func Set(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	configuration = *Parse(configFile)
}

func Parse(r io.Reader) *Config {
	conf := Config{}
	dec := yaml.NewDecoder(r)

	if decErr := dec.Decode(&conf); decErr != nil {
		return nil
	}

	return &conf
}
