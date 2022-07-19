package main

import (
	"fmt"
)

func main() {
	go greeter("Hello")
	go greeter("World")
}

//to create a go routine we use the keyword go
func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}

}
