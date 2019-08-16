package main

import (
	"fmt"

	"github.com/IguteChung/go-errors"
)

func main() {
	err := foo()
	fmt.Println(err.(errors.ErrorTracer).Stack().Format(errors.GoLikeFormatter))
}

func foo() error {
	err := bar()
	return errors.Errorf("failed to call bar: %v", err)
}

func bar() error {
	return errors.New("something wrong")
}
