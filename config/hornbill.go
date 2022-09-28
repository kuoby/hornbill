package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Hornbill struct {
	Mongoose    Mongoose    `toml:"mongoose"`
	Environment Environment `toml:"environment"`
}

type Mongoose struct {
	BaseURL string `toml:"base_url"`
}

type Environment struct {
	Provisioner string `toml:"provisioner"`
}

// ParseFile parses a file, and returns Hornbill.
func ParseFile(p string) (*Hornbill, error) {
	cfg := &Hornbill{}
	if _, err := toml.DecodeFile(filepath.Clean(p), &cfg); err != nil {
		return nil, errors.Wrap(err, "could not parse configuration from file")
	}

	return cfg, nil
}
