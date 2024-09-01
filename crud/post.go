package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//define struct

type Todo struct {
	UserId    int    `json:"user_id"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func postData(postUrl string) (string, error) {
	todo := Todo{
		UserId:    24,
		Title:     "Sambit Code",
		Completed: true,
	}

	//convert todo struct to json
	jsonData, error := json.Marshal(todo)
	if error != nil {
		fmt.Println("Error while converting to Json format:--", error)
		return "", error
	}

	//convert json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	response, error := http.Post(postUrl, "application/json", jsonReader)
	if error != nil {
		fmt.Println("Error while posting data:--", error)
		return "", error
	}
	defer response.Body.Close()

	fmt.Println("The respone status is:", response.Status)

	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("Post data response:----", data)
	return string(data), nil
}
