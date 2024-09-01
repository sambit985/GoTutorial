package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is crud api learning file")
	getUrl := "https://jsonplaceholder.typicode.com/todos/1"
	postUrl := "https://jsonplaceholder.typicode.com/posts"

	//Calling getData()
	data, err := getData(getUrl)
	if err != nil {
		fmt.Println("Error while calling get function:--", err)
		return
	}
	fmt.Println("Get data:--", data)

	//Calling postData()
	res, error := postData(postUrl)
	if error != nil {
		fmt.Println("Error while calling postData function", error)
		return
	}
	fmt.Println("Post Data result is:", res)
}
