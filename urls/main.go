package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://sam.tech:3000/contact?email=sambit@gmail.com"

func main() {
	fmt.Println("This is Url handling Code file")
	fmt.Println("My url is:-", myUrl)

	//Parsing
	result, error := url.Parse(myUrl)
	if error != nil {
		panic(error)
	}
	fmt.Println("Result:--", result)
	fmt.Println("Url scheme:-", result.Scheme)
	fmt.Println("Host :-", result.Host)
	fmt.Println("Porst :-", result.Port())
	fmt.Println("Raw query:-", result.RawQuery)

	//construct url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "dev.tech:4000",
		Path:   "/about",
	}

	finalUrl := partsOfUrl.String()

	fmt.Println("The final constructed url:-", finalUrl)
}
