package main
import (
	"fmt"
)
func main(){
	for row:=1;row<=5;row++{
		for hashNum:=1;hashNum<=6-row;hashNum++{
			fmt.Print("#")
		}
		fmt.Println()
	}
}