package main

import (
	"fmt"
)

func _myMessage() {
	fmt.Println("I am so excited")
}
func main() {
	_myMessage()
	_myMessage()
	familyname("john")
	familyname("james")
	familyname("anjna")
	familygreeting("John", 4)

}

//parameterized functions
func familyname(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

//2 argument function
func familygreeting(fname string, age int) {
	fmt.Println("Hello", age, "Years old ", fname, "Refsnes")

}
