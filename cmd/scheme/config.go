package main

type Config struct {
	Library struct {
		Path string
	}
}

func config() (cfg *Config, err error) {
	cfg = &Config{}
	// apply default values
	// apply environment values
	// apply command line values
	return nil, nil
}

