package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getData(getUrl string) (string, error) {
	res, error := http.Get(getUrl)
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
	//fmt.Println("The body is:", body)
	return string(body), nil
}
