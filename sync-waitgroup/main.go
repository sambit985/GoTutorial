package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //signal gives that worker task done
	fmt.Printf("Worker %d started\n", i)
	fmt.Printf("Wroker %d ended \n", i)
}

func main() {
	fmt.Println("This is sync.waitgroup code file")

	//create a variable for sync.waitgroup
	var wg sync.WaitGroup
	//execute 3 worker
	for i := 1; i <= 3; i++ {
		// fmt.Println(i)
		wg.Add(1) //Increment the waitgroup
		go worker(i, &wg)
	}
	//wait untill waitgroup counter becomes 0
	wg.Wait()
	fmt.Println("All worker task has finished")
}
