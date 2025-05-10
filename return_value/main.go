package main

import (
	"errors"
	"fmt"
)

func f() (rerr error) {
	defer func() {
		if r := recover(); r != nil {
			if err, isErr := r.(error); isErr {
				rerr = err
			} else {
				rerr = fmt.Errorf("%v", r)
			}
		}
	}()
	panic(errors.New("error"))
	return nil
}

func main() {
	if err := f(); err != nil {
		fmt.Println(err)
	}
}