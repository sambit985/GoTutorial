package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func UpdateTodo(updateUrl string) (string, error) {
	fmt.Println("This is the update method code")
	todo := Todo{
		UserId:    6925803,
		Title:     "Sambit is here no need to fear",
		Completed: false,
	}

	//convert todo to json
	jsonData, error := json.Marshal(todo)
	if error != nil {
		fmt.Println("Error while converting to json:--", error)
		return "", error
	}

	//convert json data to string
	jsonString := string(jsonData)

	//read json string
	jsonReader := strings.NewReader(jsonString)

	//create put request
	req, err := http.NewRequest("PUT", updateUrl, jsonReader)
	if err != nil {
		fmt.Println("Error while create request:--", err)
		return "", err
	}
	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	response, er := client.Do(req)
	if er != nil {
		fmt.Println("Error while sending request:--", er)
		return "", er
	}

	defer response.Body.Close()

	//Read http response
	resData, errr := ioutil.ReadAll(response.Body)
	if errr != nil {
		fmt.Println("Error while read http response:--", errr)
		return "", errr
	}
	//fmt.Println("Final response Data :---", string(resData))
	return string(resData), nil
}
