package main

import "fmt"

// function to operation with value which stored in perticular memory space
func ModifyValueByReference(num *int) {
	//i want to modify the num adress contained value which is 5
	*num = *num + 5 //*num used to get num adress contained data
}

func main() {
	//declare a variable
	var num = 2
	// this way we can define aswell- var ptr = &num
	//another way to declare pointer which store int type
	var ptr *int
	ptr = &num

	fmt.Println("Pointer contains:---", ptr)
	//check what data contained by pointer
	fmt.Println("Data contained by pointer:--", *ptr)

	//Example-2
	var pointer *string
	//check without assign data what adress pointer does contain
	fmt.Println("Pointer contains:--", pointer) //nil

	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	//Example-3
	//declare a variable and assign value 5
	value := 5
	//Pass the value address as arg
	ModifyValueByReference(&value)
	fmt.Println("Value modified:--", value)
}
