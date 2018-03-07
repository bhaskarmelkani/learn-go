package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addShort(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(addShort(1, 3))
}
