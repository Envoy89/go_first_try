package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Println("-> ")
		_, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}

func f(cnt *int) {
	*cnt += 1
}
