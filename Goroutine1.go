//This is code from geeksforgeeks
package main

import (
	"fmt"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		fmt.Println(str)
	}
}
func main() {
	//calling goroutine
	go display("Welecome")
	//calling normal function
	display("GeekForGeeks")
}
