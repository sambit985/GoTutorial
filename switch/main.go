package main

import "fmt"

func main() {
	day := 2

	//execute corresponding case by checking the value of day
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Unknown day")
	}

	const month = 5
	//execute corresponding case by checking the value of month
	switch month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 5:
		fmt.Println("May")
	}
}
