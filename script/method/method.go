package main

import "fmt"

type (
	Hex int
	T   int
)

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func (t *T) f() { println("hi") }

func main() {
	var hex Hex = 100
	// メソッド値
	// f := hex.String
	// fmt.Println(f())
	// メソッド式
	f := Hex.String
	fmt.Printf("%T\n%s\n", f, f(hex))
	fmt.Println(hex.String())

	var v T
	(&v).f()
	v.f()
}
