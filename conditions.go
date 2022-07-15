package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 5
	//greater than operator
	fmt.Println(x > y)

	//not equal to operator !=
	fmt.Println(x != y)

	//logical and
	fmt.Println((x > y) && (x < y))
	//logical or operator
	fmt.Println((x > y) || (x < y))

	//using if to check the condition
	//if is use to print the block of code if the specified condtion is true
	// we can apply test the conditon either by variable or by the numbers
	if x > y {
		fmt.Println("Yes the x is greater")
	}
	//using if else to check the condition
	// else is use to print the block of code if the specified condition is false
	//we can apply the if else condition on eiher numbers or variables
	if x < y {
		fmt.Println("Yes x is lesser than y ")
	} else {
		fmt.Println("No is not lesser than y")
	}
	//else if is use to specify the second condition if the first condition is false
	time := 22
	if time < 10 {
		fmt.Println("Good Morning")
	} else if time < 20 {
		fmt.Println("Good Day")
	} else {
		fmt.Println("Good evening")
	}

	//nested if
	//in nested if we specify if inside if
	num := 20
	if num > 10 {
		fmt.Println("Numbers a greter than 10")
		if num > 15 {
			fmt.Println("Numbers are also greater than 15")
		}
	} else {
		fmt.Println("Nubmbers are lesser than 10")
	}

}
