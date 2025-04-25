package main

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
}

