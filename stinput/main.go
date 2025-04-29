package main

import (
	"fmt"
	"log"
	"os"
)

func f() error {
	return nil
}

func main() {
	fmt.Fprintln(os.Stderr, "エラー")
	fmt.Fprintln(os.Stdin, "入力")
	fmt.Fprintln(os.Stdout, "Hello,")
	os.Exit(1)

	if err := f(); err != nil {
		log.Fatal(err)
	}
}
