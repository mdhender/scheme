package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(2)
	}
}

func run() error {
	return repl([]byte{})
}

func repl(b []byte) error {
	return nil
}

type sexp struct {}


func read(r io.ByteReader) (*sexp, error) {
	return nil, fmt.Errorf("!implemented")
}