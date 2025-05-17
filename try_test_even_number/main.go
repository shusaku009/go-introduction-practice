package main

import (
	"bufio"
	"fmt"
	"os"
)

func isEven(n int) bool {
	return n % 2 == 0
}

func proccessInput(input string) (bool, error) {
	var num int
	_, err := fmt.Sscanf(input, "%d", &num)
	if err != nil {
		return false, err
	}
	return isEven(num), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result, err := proccessInput(scanner.Text())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}