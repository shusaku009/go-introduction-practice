package main

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
}