package main

import "fmt"

type Hex int

type Stringer interface {
	String() string
}

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type Hex2 struct{ Hex }

func main() {
	var s Stringer
	h := Hex(100)
	s = h
	fmt.Println(s.String())

	h2 := Hex2{h}
	s = h2
	fmt.Println(s.String())
}
