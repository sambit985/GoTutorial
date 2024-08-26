package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getData() string {
	res, error := http.Get("https://api.restful-api.dev/objects")
	if error != nil {
		fmt.Println("Error while getting data:-", error)
		return ""
	}
	defer res.Body.Close()
	//read res
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error while reading data", err)
		return ""
	}
	//fmt.Println("The res data is:", string(data))
	return string(data)

}

func main() {
	fmt.Println("This is web request code file")
	//call getdata function and store return response
	data := getData()
	fmt.Println("Data recieved:--", data)

}
