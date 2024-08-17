package main

import "fmt"

// Create normal function with func keyword
func simpleFunction() {
	fmt.Println("Simple Function is calling")
}

// Create function for add 2 numbers
func add(a int, b int) int { //int which declare in outside() is for return type
	return a + b
}

// create function for multiply 2 numbers
func mult(a, b int) int {
	var result int = a * b //or we can write "result := a*b"
	fmt.Println("Result is", result)
	return result
}

// Main function is executing by default
func main() {
	fmt.Println("Main function is executing")
	simpleFunction()
	ans := add(4, 4) //store the return res into the ans variable
	fmt.Println("Addition of 2 number is:--", ans)

	ans2 := mult(2, 4)
	fmt.Println("Multiplication of 2 number is:--", ans2)

}
