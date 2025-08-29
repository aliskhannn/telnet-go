package config

import "time"

// Config represents the configuration options for the telnet client.
type Config struct {
	Timeout time.Duration // connection timeout
}

// New creates and returns a new Config instance with the given parameters.
func New(timeout time.Duration) *Config {
	return &Config{
		Timeout: timeout,
	}
}
