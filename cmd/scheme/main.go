package main

import (
	"fmt"
	"github.com/mdhender/scheme"
	"io"
	"os"
)

func main() {
	i := scheme.New()
	fmt.Println(*i)

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

func repl(b []byte) error {
	return nil
}

type sexp struct {}


func read(r io.ByteReader) (*sexp, error) {
	return nil, fmt.Errorf("!implemented")
}