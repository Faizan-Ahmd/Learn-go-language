//Go has different kind of values like int, float , string and boolean etc.
//here few examples.
//string can be added through + operation. Intergers and float.
//Boolean with the boolean operation as you have expect..
package main
import (
	"fmt"
)
func main(){
	fmt.Println("go"+" Lang")
	fmt.Println("1+1 = ",1+1)
	fmt.Println("7.0/3.0 =",7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}