package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welecome! to the main function")
	//creating Anonymous Goroutine
	go func() {
		fmt.Println("Welcome to the InvoZone")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Good bye! to the main function")
}
