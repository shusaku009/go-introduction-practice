package main

import (
	"errors"
	"fmt"
)

type MyError struct{ No int }

func (err *MyError) Error() string {
	return fmt.Sprintf("Error %d", err.No)
}

func main() {
	err1 := &MyError{No: 100}
	err2 := errors.New("err2")
	err := fmt.Errorf("%w and %w", err1, err2)

	var myerr *MyError
	if errors.As(err, &myerr) {
		fmt.Println(myerr)
	}
}
