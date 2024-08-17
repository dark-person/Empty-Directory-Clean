package config

import (
	"fmt"
	"os"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Global koanf instance. Use "." as the key path delimiter.
var k = koanf.New(".")

// Load yaml config from given path.
//
// If failed to load config, then a default config will be returned.
func LoadYaml(path string) (*RemovalConfig, error) {
	// Check if file exist
	if _, err := os.Stat(path); err != nil {
		return Default(), fmt.Errorf("path %s does not exist", path)
	}

	// Start Load file
	err := k.Load(file.Provider(path), yaml.Parser())
	if err != nil {
		return Default(), err
	}

	// Load config to struct
	c := &RemovalConfig{
		Files:     k.Strings("files"),
		Extension: k.Strings("extensions"),
	}

	return c, nil
}
