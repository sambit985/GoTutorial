package main

import (
	"encoding/json"
	"fmt"
)

// create struct
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("This is Json learning code")
	// var person Person
	// person.Name = "Sambit"
	// fmt.Println(person)

	person := Person{
		Name:    "Sambit",
		Age:     24,
		IsAdult: true,
	}
	//fmt.Println("person is :-", person)

	//convert person into JSON using Marshalling
	jsonData, error := json.Marshal(person)
	if error != nil {
		panic(error)
	}

	fmt.Println("Encoded data:--", jsonData)

	//convert or decode encoded json byte[] data to json and store in same place

	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		panic(err)
	}
	fmt.Println("Person data after decode:--", person)

}
