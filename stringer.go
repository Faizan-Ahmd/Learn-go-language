package main

import "fmt"

type person struct {
	id   int
	Name string
	Age  int
}

func main() {
	var I = person{id: 1, Name: "Logan", Age: 22}
	fmt.Println(I)
}
