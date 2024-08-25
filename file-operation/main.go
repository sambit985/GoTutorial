package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("This is File operation code file")

	//create file
	/* file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file:--", err)
		return
	}

	fmt.Println("The file successfully created:", file)
	//In last we have clear the resource
	defer file.Close()

	//write content in the file using WriteString() which return 2 values i.e-slice of byte and error
	_, error := io.WriteString(file, "I am Sambit")
	if error != nil {
		fmt.Println("Error while writing file:--", error)
	}
	fmt.Println("Successfully write inside the example.txt file")
	*/

	//Open the file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening the file:-", err)
		return
	}

	defer file.Close()

	// first create buffer
	buffer := make([]byte, 1024)

	//Read the file and store the buffer data to buffer slice
	//Infinite for loop
	for {
		//Read file which giving chunk data or byte data, at the same time  push to buffer chunk wise
		_, error := file.Read(buffer)
		//when it reach the end of file
		if error == io.EOF {
			break
		}
		if error != nil {
			fmt.Println("Error while reading file:--", error)
			return
		}
	}
	//process the buffer read by converting to string
	fmt.Println(string(buffer))

	//ioutil.REadFile() or os.ReadFile() we can use for same purpose but it is not good approach when reading big file beacuse it is taking whole data and puting into memory at a single time

	content, errors := os.ReadFile("example.txt")
	if errors != nil {
		fmt.Println("Error while reading file")
		return
	}
	fmt.Println(string(content))

}
