package main

import (
	"fmt"
)

func main() {
	fmt.Println("Channels in go lang")
	mych := make(chan int)
	mych <- 5
	fmt.Println(<-mych)
}
