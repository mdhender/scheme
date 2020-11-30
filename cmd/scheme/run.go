package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/mdhender/scheme"
)

func run(cfg *Config) error {
	var cli struct {
		Debug bool `help:"Enable debug mode."`
		Eval EvalCmd `cmd help:"Evaluate scripts."`
	}
	ctx := kong.Parse(&cli)

	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
	return nil
}

type Context struct {
	Debug bool
}

type EvalCmd struct {
	Paths []string `arg optional name:"path" help:"Scripts to evaluate." type:"path"`
}

func (cmd *EvalCmd) Run(ctx *Context) error {
	i := scheme.New()
	for _, fname := range cmd.Paths {
		expr := fmt.Sprintf("(load %q)", fname)
		fmt.Printf("(eval %s)\n", expr)
	}
	fmt.Println(*i)
	return nil
}
