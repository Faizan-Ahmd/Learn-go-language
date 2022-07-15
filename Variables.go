package main

import (
	"fmt"
)

func main() {
	var username string = "Faizan"
	fmt.Println(username)
	fmt.Printf("the variable is a type of : %T \n", username)

	var isOn bool = true
	fmt.Printf("The Value of is on is %v\n", isOn)
	fmt.Printf("The type of is on is %T\n", isOn)

	var price int = 23
	fmt.Printf("The value of price is %v \n", price)
	fmt.Printf("The type of price id %T\n", price)

	var smallInt uint = 255
	fmt.Printf("The value of smallInt is %v\n", smallInt)
	fmt.Printf("The type of smallInt is %T\n", smallInt)
}
