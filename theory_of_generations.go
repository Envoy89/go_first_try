package main

import "fmt"

func main() {
	bornYear := 1994

	switch {
	case bornYear >= 1946 && bornYear <= 1964:
		fmt.Println("Привет, бумер!")
	case bornYear >= 1965 && bornYear <= 1980:
		fmt.Println("Привет, представитель X!")
	case bornYear >= 1981 && bornYear <= 1996:
		fmt.Println("Привет, миллениал!")
	case bornYear >= 1997 && bornYear <= 2012:
		fmt.Println("Привет, зумер!")
	case bornYear >= 2013:
		fmt.Println("Привет, альфа!")
	}
}
