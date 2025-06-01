package main

import (
	"fmt"
	"sync"
)

func f() { fmt.Println("f") }

func main() {
	var once sync.Once
	once.Do(f)
	once.Do(f)
	fmt.Println("done")
}