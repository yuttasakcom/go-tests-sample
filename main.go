package main

import "fmt"

func main() {
	sum := add(1, 2)

	fmt.Println(sum)
}

func add(lhs, rhs int) int {
	return lhs + rhs
}
