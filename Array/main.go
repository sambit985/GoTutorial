package main

import "fmt"

func main() {
	fmt.Println("Learning Array")
	//create an array
	var names [4]string
	//add elements to the array
	names[0] = "Radha"
	names[1] = "Sambit"
	//Print names array
	fmt.Println("Names:--", names)
	//Check total size of an array
	fmt.Println("Names Array Size :--", len(names))
	names[3] = "Ram"
	fmt.Println("Names:---", names)

	//Initialize array with elements
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:--", numbers)
	//Check length of an array
	fmt.Println("Length of Numbers:--", len(numbers))
	//Acess element from an array
	fmt.Println("Value of second index:--", numbers[2])

	//Initialize an empty arrays
	var price [5]float32
	fmt.Println("Price array:--", price)
}
