package main

import (
	"os"
)

type Config struct {
	host string
	port string
}

func (c *Config) Address() string {
	return c.host + ":" + c.port
}

func newConfig() Config {
	c := Config{}
	c.host = getOr("HOST", "0.0.0.0")
	c.port = getOr("PORT", "8125")
	return c
}

func getOr(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}
