package main

import "fmt"

type Func func() string

func (f Func) String() string { return f() }

var _ fmt.Stringer = Func(nil)

func main() {
	ns := []int{1, 2, 3, 4}
	// できない
	var vs []any = ns
}
