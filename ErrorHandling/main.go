package main

import "fmt"

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("dominator must not be 0") //return float type and error type
	}
	return a / b, nil //nil indicates not error
}
func main() {
	ans, err := division(10, 0)
	//if statement must be a boolean expression. The error type in Go is not a boolean; it's an interface type that can either hold a value (indicating an error occurred) or be nil (indicating no error).
	if err != nil { //if err != not indicates error
		fmt.Println("Error handling")
	}
	fmt.Println("Division of 2 Number:--", ans)
}
