package main

import (
	"fmt"

	"github.com/IguteChung/go-errors"
)

func init() {
	errors.ApplyFormatter(errors.GoLikeFormatter)
}

func main() {
	err := foo()
	fmt.Println(errors.StackTrace(err))
}

func foo() error {
	err := bar()
	return errors.Errorf("failed to call bar: %v", err)
}

func bar() error {
	return errors.New("something wrong")
}
