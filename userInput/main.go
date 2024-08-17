package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, What's Your Name?")
	var name string
	//read user input untill empty sapce
	fmt.Scan(&name)
	fmt.Println("Hello Mr.", name)

	//Read all full linr user input untill new line
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n') //used "_" to ignore error or other value
	fmt.Println("Hello Mr.", name)
}
