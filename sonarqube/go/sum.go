package main

import "fmt"

func main() {
	fmt.Println(Sum(15, 15))
}

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
