package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getData(url string) (string, error) {
	res, error := http.Get(url)
	if error != nil {
		fmt.Println("Error while frtching data:--", error)
		return "", error
	}

	defer res.Body.Close()

	//read res.Body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	fmt.Println("This is crud api learning file")
	url := "https://jsonplaceholder.typicode.com/todos/1/"

	data, err := getData(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Get data:--", data)
}
