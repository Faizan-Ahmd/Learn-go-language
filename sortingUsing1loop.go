package main

import "fmt"

type sliceint []int

func Sorting(si sliceint, length int) *sliceint {
	for i := 0; i < length-1; i++ {
		if si[i] > si[i+1] {
			temp := si[i]
			si[i] = si[i+1]
			si[i+1] = temp
			i = -1
		}

	}
	return &si
}
func main() {
	var slice = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	si := (Sorting(slice, len(slice)))
	//for i := 0; i < len(slice); i++ {
	fmt.Println(*si)
	//	}
}
