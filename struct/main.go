package main

import "fmt"

// define struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	City    string
	State   string
	ZipCode int
}

type Contact struct {
	Email string
	Phone int
}

type Employee struct {
	Person_Details  Person
	Person_Address  Address
	Person_Contacts Contact
}

func main() {
	//1st way to use by creating an instance of struct 'Person'
	var sambit Person
	fmt.Println("Sambit person:-", sambit)
	sambit.Age = 24
	sambit.FirstName = "Sambit"
	sambit.LastName = "Nayak"
	fmt.Println("Sambit:--", sambit)

	//2nd way to use by initializing struct literal
	person1 := Person{
		FirstName: "Sohan",
		LastName:  "Singh",
		Age:       25,
	}
	fmt.Println("Person1:--", person1)

	//3rd way
	var person2 = new(Person)
	person2.FirstName = "Rajesh"
	person2.LastName = "Srhibasta"
	person2.Age = 26

	fmt.Println("Person2:--", person2)

	//Nested Struct
	employee := Employee{
		Person_Details: Person{
			FirstName: "Mahatma",
			LastName:  "Gandhi",
			Age:       154,
		},
		Person_Contacts: Contact{
			Email: "mahatma@gmail.com",
			Phone: 7349673455,
		},
		Person_Address: Address{
			City:    "Porbandar",
			State:   "Gujurat",
			ZipCode: 459430,
		},
	}

	fmt.Println("Employee:-", employee)
	fmt.Println("Employee Phone:-", employee.Person_Contacts.Phone)

}
