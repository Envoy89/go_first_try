package main

import "fmt"

type WeekDay int

const (
	Monday WeekDay = iota + 1
	Tuesday
	Wednesdey
	Thursday
	Friday
	Saturday
	Sunday
)

func NextDay(day WeekDay) WeekDay {
	return (day % 7) + 1
}

func main() {
	var today WeekDay = Monday
	tommorow := NextDay(today)
	fmt.Println("today =", today, "tommorow =", tommorow)
	fmt.Println(today % 7)
}
