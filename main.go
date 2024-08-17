package main

import (
	"fmt"
	//custompackage "sambitTutorial/myUtility" //import custom module
)

func main(){
	//fmt.Println("Hello, World");
	//custompackage.PrintMessage("Execute by another module"); //use custom module function to execute 
	//Variable declaration
	const name string = "Sambit" //can use data type
	var person  = "Sambit Nayak" //auto detech data type
	var version = 1.0 //without datatype detect automatically
	 version = 34  //update 10. -> 34 due to var type

	fmt.Println("Name:--", name);
	fmt.Println("version:--", version);
	fmt.Println("Person:--", person);


	designation := "Software Engineer" //without variable type and data type
	fmt.Println("Designation:--", designation);
	
	Public := "Data is Published" //If variable name start with capital letter then dirctly use in another file as export variable
	private := "Data is Private"
	fmt.Println(Public)
	fmt.Println(private)
}