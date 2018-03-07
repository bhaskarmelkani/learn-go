package main

import "fmt"
import "math/rand"

func main() {
	rand.Seed(2)
	fmt.Println("The number is", rand.Int())
}
