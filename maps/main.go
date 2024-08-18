package main

import "fmt"

func main() {
	//create map with name-age pair using make()
	// person := make(map[string]int)
	// person["sambit"] = 24
	// person["Rakesh"] = 28

	//Create map with name-age pair initial value
	person := map[string]int{
		"sambit": 24,
		"rakesh": 28,
	}

	fmt.Println("Person:--", person)
	//Access age of sambit
	fmt.Println("Sambit age is:", person["sambit"])
	//Modify Rakesh age
	person["rakesh"] = 27
	fmt.Println("Rakesh is:", person["rakesh"])

	//retrive data using for loop with range
	for name, age := range person {
		fmt.Printf("Name is: %s and age is: %d\n", name, age)
	}

	//checking if exist or not
	value, exists := person["dhoni"] //map return value with boolean(true/false)
	if exists {
		fmt.Println("Age of Dhoni is:", value)
	}

	//delete an element from person map using delete(map,key)
	delete(person, "rakesh")
	fmt.Println("Person:--", person)
}
