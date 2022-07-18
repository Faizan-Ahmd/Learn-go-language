package main

import (
	"fmt"
)

func PrintAnyThing[T any](v T) T {
	return v
}
func main() {
	x := 4
	y := "HelloWorld"
	fmt.Println(PrintAnyThing(x))
	fmt.Println(PrintAnyThing(y))
}
