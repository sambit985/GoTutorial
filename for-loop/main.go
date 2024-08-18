package main

import "fmt"

func main() {

	//Example-1
	//initialize i to 0 and check if i<5 and increment i++ after each iteration
	for i := 0; i < 5; i++ {
		fmt.Println("Current value:-", i)
	}

	//Example-2
	//Infinite loop
	count := 0
	for {
		fmt.Println("counter:-", count)
		//increment 1 every time
		count++
		//can add condition aswell when to stop infinite loop
		if count == 10 {
			break
		}
	}

	//Example-3
	//range:- Idiomatic way to loop through each element
	numbers := []int{19, 23, 35, 40, 52}
	//Iterate over array/slice in general way
	for i := 0; i < len(numbers); i++ {
		//fmt.Println("Current Element:-", numbers[i])
		fmt.Printf("Index is %d, value is %d\n", i, numbers[i])
	}
	//Iterate over array/slice using range
	for index, value := range numbers {
		fmt.Printf("Index is: %d, value is: %d\n", index, value)
	}

	//Example-5
	//Iterate over a string using range
	name := "Sambit is a backend developer"
	for index, char := range name {
		fmt.Printf("Index is: %d, Charecter is: %c\n", index, char)
	}

	//Example-6
	//ignore the index in range response
	for _, value := range numbers {
		fmt.Println("Number is:", value)
	}

}
