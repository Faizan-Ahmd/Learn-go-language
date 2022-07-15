package main

import (
	"fmt"
)

func main() {
	var value interface{} = "GeeksforGeeks"
	var value1 string = value.(string)
	fmt.Println(value1)
	var value2 int = value.(int)
	fmt.Println(value2)
}
