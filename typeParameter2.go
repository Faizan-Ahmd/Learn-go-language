package main

import (
	"fmt"
)

func Any[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}
func main() {
	s := []int{1, 23, 3, 34, 55}
	ss := []string{"Hello", "World", "How", "Are", "You"}
	Any(s)
	Any(ss)
}
