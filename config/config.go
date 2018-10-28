package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server ServerInfo
}

type ServerInfo struct {
	Addr     string
	Dir      string
	ValueDir string
	Debug    string
}

var DefaultConfig Config

func NewConfig(fileName string) (*Config, error) {
	if _, err := toml.DecodeFile(fileName, &DefaultConfig); err != nil {
		return nil, err
	}
	return &DefaultConfig, nil
}
