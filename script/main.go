package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, 世界\n")
	fmt.Printf("%d-%s", 100, "偶数")

	var price int
	fmt.Print("値段>")
	fmt.Scanln(&price)
	fmt.Printf("%d円\n", price)

	msg := "hello"

	fmt.Println(msg)
	fmt.Println("世" + "界")

	hello := "Hello, 世界"
	fmt.Println(hello)

	const helloWorld = "Hello, 世界"
	fmt.Println(helloWorld)

	const (
		a = 1 + 2
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		StatusOK = iota
		StatusNG
	)

	const (
		FlagA = 1 << iota
		FlagB
		FlagC
		FlagD
	)

	n := 100 + 200
	m := n + 100
	msg2 := "hoge" + "fuga"

	fmt.Println(n, m, msg2)

	for i := 1; i <= 100; i = i + 1 {
		if i%2 == 1 {
			fmt.Println(i, "- 奇数")
		} else {
			fmt.Println(i, "- 偶数")
		}
	}

	for i := 1; i <= 100; i = i + 1 {
		switch i % 2 {
		case 1:
			fmt.Println(i, "- 奇数")
		default:
			fmt.Println(i, "- 偶数")
		}
	}
}
