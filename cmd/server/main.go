package main

import (
	"fmt"
	"os"
)

func main() {
	cfg, err := config()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}
	if err := run(cfg); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}
}

func run(cfg *Config) error {
	return fmt.Errorf("!implemented")
}

