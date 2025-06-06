package main

import "fmt"

func makeCh() chan int {
	return make(chan int)
}

func recvCh(recv chan int) int {
	return <-recv
}

func main() {
	ch := makeCh()
	go func() {ch <- 100 }()
	fmt.Println(recvCh(ch))
}