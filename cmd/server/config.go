package main

import (
	"fmt"
	"github.com/mdhender/scheme/pkg/xii"
	"log"
)

type Config struct {
	Server struct {
		Scheme string
		Host string
		Port int
	}
}

// config returns an initialized configuration.
// It applies default values, then values from the environment, then from the command line.
// In effect, that means that command line overrides environment overrides default values.
func config() (cfg *Config, err error) {
	cfg = &Config{}
	// set default values
	cfg.Server.Scheme = "http"
	cfg.Server.Host = "localhost"
	cfg.Server.Port = 8080
	// get environment values
	key := "SERVER_SCHEME"
	if cfg.Server.Scheme, err = xii.AsString(key, xii.StringOpts{DefaultValue: cfg.Server.Scheme}); err != nil {
		return nil, err
	}
	log.Printf("[config] env %-30s == %q\n", key, cfg.Server.Scheme)
	if !(cfg.Server.Scheme == "http" || cfg.Server.Scheme == "https") {
		return nil, fmt.Errorf("invalid scheme %q", cfg.Server.Scheme)
	}
	key = "SERVER_HOST"
	if cfg.Server.Host, err = xii.AsString(key, xii.StringOpts{DefaultValue: cfg.Server.Host}); err != nil {
		return nil, err
	}
	log.Printf("[config] env %-30s == %q\n", key, cfg.Server.Host)
	key = "SERVER_PORT"
	if cfg.Server.Port, err = xii.AsInt(key, xii.IntOpts{DefaultValue: cfg.Server.Port}); err != nil {
		return nil, err
	}
	log.Printf("[config] env %-30s == %d\n", key, cfg.Server.Port)

	// get command line values
	return cfg, nil
}
