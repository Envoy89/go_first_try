package main

import "fmt"

func main() {
	var str string

	if str == "" {
		fmt.Println("Empty!")
	}
	if len(str) == 0 {
		fmt.Println("Empty!")
	}

	var i int
	var i32 int32

	println(int32(i) == i32)

}
