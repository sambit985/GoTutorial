package main

import "fmt"

func main() {
	age := 20
	if age > 18 {
		fmt.Println("You are eligible for voting")
	} else {
		fmt.Println("You are not eligible for voting")
	}

	//Find big among 3
	var a, b, c = 13, 20, 18
	if a > b && a > c {
		fmt.Println("a is biggest:", a)
	} else if b > c && b > a {
		fmt.Println("b is biggest:", b)
	} else {
		fmt.Println("C is biggest:", c)
	}
}
