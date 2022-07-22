package main

import (
	"fmt"
)

func main() {
	go greeter("Hello")
	go greeter("World")
}

//to create a go routine we use the keyword go

//when we are using go function it only prints the non go function the reason is that
//you did not mentioned when to come on go function
func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}

}
