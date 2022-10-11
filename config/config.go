package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Configuration struct {
	Mongoose    Mongoose    `toml:"mongoose"`
	Environment Environment `toml:"environment"`
}

// ParseFile parses a file, and returns Configuration.
func ParseFile(p string) (*Configuration, error) {
	cfg := &Configuration{}
	if _, err := toml.DecodeFile(filepath.Clean(p), &cfg); err != nil {
		return nil, errors.Wrap(err, "could not parse configuration from file")
	}
	return cfg, nil
}
