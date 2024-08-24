package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello This is data conversion code")

	var number = 8
	fmt.Println("THe number is:--", number)
	fmt.Printf("The number is %T\n", number)

	//convert int data type to float
	var data = float64(number)
	fmt.Println("The data is:", data)
	fmt.Printf("The type of data is : %T\n", data)

	//Convert int to string
	num := 123
	var num_str = strconv.Itoa(num)
	fmt.Println("The num_str :--", num_str)
	fmt.Printf("The num Type is:-- %T\n", num_str)

	//convert string to int
	var numStr = "456"
	str_num, _ := strconv.Atoi(numStr)
	fmt.Println("The str_num is:--", str_num)
	fmt.Printf("The type of str_num is:-- %T\n", str_num)

	//convert string to float
	num_string := "6.78"
	float_str, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("The string is:-- ", num_string)
	fmt.Printf("The type of float_str is - %T\n", float_str)

}
