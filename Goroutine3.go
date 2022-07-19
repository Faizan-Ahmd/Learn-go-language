package main

import (
	"fmt"
	"time"
)

func display(str string) {
	fmt.Println(str)
}
func main() {
	go display("Hello1")
	go display("Hello2")
	go display("Hello3")
	go display("Hello4")
	go display("Hello5")
	go display("Hello6")
	go display("Hello7")
	go display("Hello8")
	go display("Hello9")
	go display("Hello10")
	go display("Hello11")
	display("World")
	time.Sleep(1 * time.Millisecond)
}
