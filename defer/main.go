package main

import "fmt"

func main() {
	//defer used to call in last
	defer fmt.Println("This is defer code file")
	defer fmt.Println("This is 1st line")
	fmt.Println("This is last line")

	newFunction()
}

func newFunction() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("The number is:-", i)
	}
}

//this is defer code file, this is 1st line, 0, 1, 2, 3, 4   - stored in stack

// 4,3,2,1,0,this is 1st line, this is defercode  - LIFO

//This is last line, 4,3,2,1,0,this is 1st line, this is defercode  - result
