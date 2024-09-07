package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func deleteTodo(deleteUrl string) string {
	fmt.Println("This is delete http method code")
	//create delete request
	req, err := http.NewRequest(http.MethodDelete, deleteUrl, nil)
	if err != nil {
		fmt.Println("Error while create delete request:", err)
		return ""
	}
	//send the req through client
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending request through client:", err)
		return ""
	}
	defer response.Body.Close()

	//read http response
	byteData, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("Error while reading reading http response:--", error)
		return ""
	}
	//check status code
	fmt.Println("The delete response status:", response.Status)
	//convert byteData to string and return
	return string(byteData)
}
