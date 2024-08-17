package main

import "fmt"

func main() {
	//Declare variables
	age := 25
	height := 8.234656
	name := "Alice"

	//Formatted print
	fmt.Printf("Name is %s \n", name)
	fmt.Printf("Age is %d \n", age)
	fmt.Printf("Height is %.3f \n", height)

	//Check data type of variable
	fmt.Printf("Type of Age is %T \n", age)
	fmt.Printf("Type of name is %T\n", name)
}
