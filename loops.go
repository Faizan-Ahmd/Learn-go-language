package main

import (
	"fmt"
)

func main() {
	adj := [2]string{"Big", "Tasty"}
	fruit := [3]string{"Apple", "Orange", "Mango"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruit); j++ {
			fmt.Println(adj[i], fruit[j])
		}
	}

	//using range to print the indexes and the values of the array
	for index, val := range fruit {
		if index == 1 {
			break
		}
		fmt.Printf("%v\t%v\n", index, val)
	}
}
