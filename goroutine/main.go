package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, wrold")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Wait for 2 second")
}

func sayHi() {
	fmt.Println("Hii Sambit")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("This is Goroutine Code file")
	go sayHello()
	go sayHi()

	//wait for goroutine to finish
	time.Sleep(3000 * time.Millisecond)
}
