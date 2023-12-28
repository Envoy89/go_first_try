package main

import "fmt"

func main() {
	var a, b int

	a, b = 2, 4

	fmt.Println(a)
	fmt.Println(b)

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
