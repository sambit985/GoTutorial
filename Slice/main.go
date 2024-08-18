package main

import "fmt"

func main() {
	fmt.Println("Learning Slice")

	//Create slice with dynamic feature
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:--", numbers)
	//Add more elements to the numbers
	numbers = append(numbers, 6, 7, 8, 9)
	fmt.Println("After added Numbers:--", numbers)
	//modify the 0th index value
	numbers[0] = 13
	fmt.Println("After modify index Numbers:--", numbers)
	//check lenghth which actualnumber of elements
	fmt.Println("Length of numbers:--", len(numbers))
	//Check Capacity which allocate space
	fmt.Println("Capacity of numbers:--", cap(numbers))
	//Check data type
	fmt.Printf("Data type of numbers:--%T\n", numbers)
	//Another way to declare slice with length & capacity
	var prices = make([]float64, 3, 5)
	fmt.Println("Prices:--", prices)
	fmt.Println("Capacity of prices:--", cap(prices))
	fmt.Println("Length of the Prices:--", len(prices))

}
