package main

import (
	"fmt"
	"time"
)

func addOneLoop(number *int) {
	for {
		*number = *number + 1
		time.Sleep(time.Second)
	}
}

func printLoop(number *int) {
	for {
		fmt.Println("Test ", *number)
		time.Sleep(time.Second)
	}
}

func main() {
	/*
	The program contains two goroutines. The first goroutine adds 1 to the
	integer with a specified delay. The second goroutine prints the above
	integer also with a delay.  The goal is to print the integer while it
	increments however due to interleaving of concurrency the program will
	not print the increments in order. This is the race condition is observed.
	 */
	var number int

	go addOneLoop(&number)
	go printLoop(&number)

	time.Sleep(time.Second * 10)
	fmt.Println("END")
}
