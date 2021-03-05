package main

import "fmt"

func main() {
	fmt.Println(sum(4, 5))
}

func sum(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func multi(a int, b int) int {
	return a * b
}
