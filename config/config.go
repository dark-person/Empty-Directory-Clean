package config

import "fmt"

// Config to specific which file should removed.
type RemovalConfig struct {
	Files     []string // Filename to be removed, must be exact match
	Extension []string // Extension to be remove
}

// Get default configuration. This should be used when no config file can be loaded.
func Default() *RemovalConfig {
	return &RemovalConfig{
		Files:     []string{".DS_Store"},
		Extension: []string{},
	}
}

// String representation for removal config. Debug usage.
func (c *RemovalConfig) String() string {
	return fmt.Sprintf("{Files: %v, Extension: %v}", c.Files, c.Extension)
}
