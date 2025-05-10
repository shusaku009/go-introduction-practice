package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err := errors.Join(err1, err2)
	fmt.Println(err)
}