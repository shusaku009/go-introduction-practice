package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Score struct {
	UserID  string
	GamesID int
	Point   int
}

func main() {
	var f float64 = 10
	var n int = int(f)
	println(n)

	var sum float64
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 {
		println("good")
	}

	var a, b, c bool
	if a && b || !c {
		println("true")
	} else {
		println("false")
	}

	p := Person{
		name: "Gopher",
		age:  10,
	}
	p.age++
	println(p.name, p.age)

	msg := "Hello, World"
	func() {
		println(msg)
	}()

	fs := make([]func(), 3)
	for i := range fs {
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs {
		f()
	}

	person := struct {
		age  int
		name string
	}{age: 20, name: "Gopher"}
	p2 := person
	p2.age = 30
	p3 := person
	p3.name = "testuser"
	println(p.age, p.name)
	println(p2.age, p2.name)
	println(p3.age, p3.name)

	ns := []int{10, 20, 30}
	ns2 := ns
	ns2[1] = 50
	println(ns[0], ns[1], ns[2])
	println(ns2[0], ns2[1], ns2[2])

	for i := 1; i < 100; i++ {
		print(i)
		if isEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

func isEven(n int) bool {
	return n%2 == 0
}
