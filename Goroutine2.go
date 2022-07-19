// the program to illustrate the working of go routines
package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func main() {
	//calling go routine
	go display("Welcome")
	// calling normal function
	display("GeeksForGeeks")
}

// we made a sleep() method which makes the main go routine sleep or wait for one second
//after the waite is over the execution of main start and it perform its operations
//and then goes for the sleep the process continues until the w<6
//when the main goes for the sleep goroutine is executed
