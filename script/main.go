package main

import "fmt"

func main() {
	fmt.Printf("Hello, 世界\n")
	fmt.Printf("%d-%s", 100, "偶数")

	var price int
	fmt.Print("値段>")
	fmt.Scanln(&price)
	fmt.Printf("%d円\n", price)
}
