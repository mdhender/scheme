// Package scheme implements an R7 small interpreter.
package scheme

import "fmt"

// interpreter implements the read-eval-loop against a buffer.
type interpreter struct {

}

func New() *interpreter {
	return &interpreter{}
}

type expr struct {
	err error
}

// Read returns the next expression, the number of bytes consumed.
// If an error is encountered, it is returned as the expression.
func (i *interpreter) Read(b []byte, env *expr) (int, *expr, *expr) {
	return 0, &expr{err: fmt.Errorf("assert(read implmented)")}, env
}

// Eval evaluates the expression in the context of some environment.
// It may update the environment as a side-effect of evaluating the expression.
// It returns the result of evaluating the expression.
func (i *interpreter) Eval(e, env *expr) (*expr, *expr) {
	return &expr{err: fmt.Errorf("assert(eval implmented)")}, env
}

// Print returns the representation of the expression as a string.
func (i *interpreter) Print(e, env *expr) (string, *expr, *expr) {
	return "", &expr{err: fmt.Errorf("assert(print implmented)")}, env
}