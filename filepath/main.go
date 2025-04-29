package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("dir", "main.go")
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Dir(path))
}
