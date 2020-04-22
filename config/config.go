package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	ContentPath  string
	TemplatePath string
	StaticPath   string
	JsonPath     string
}

var config *Config

func init() {

	cp := filepath.Join(os.Getenv("GOPATH"),
		"src/github.com/rleibl/brauspielhaus/data/")

	c := Config{
		ContentPath:  cp,
		TemplatePath: filepath.Join(cp, "templates/*"),
		StaticPath:   filepath.Join(cp, "static/"),
		JsonPath:     filepath.Join(cp, "json/"),
	}
	config = &c
}

func GetConfig() *Config {
	return config
}
