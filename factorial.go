package main

import (
	"fmt"
)

func FirstFactorial(num int) int {
	/*if num > 1 {
		return FirstFactorial(num) * FirstFactorial(num-1)
	}else if (num == 0 || num == 1){
		return 1
	}*/
	fact := 1
	for i := 1; i <= num; i++ {
		fact = fact * i
	}
	return fact
}
func main() {
	fmt.Println(FirstFactorial(4))
}
