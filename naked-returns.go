package main

import (
	"fmt"
)

func naked(a, b int) (d, e int) {
	d = a
	e = b
	return
}

func main() {
	fmt.Println(naked(1, 2))
}
