package main

import (
	"fmt"
)

type person struct {
	id   int
	name string
	age  int
}

func main() {
	var i = person{id: 2, name: "jazz", age: 22}
	fmt.Println(i)
}
func (p person) String() string {
	return fmt.Sprintf("The type person \nid %d\nname %s\nage %d", p.id, p.name, p.age)
}
