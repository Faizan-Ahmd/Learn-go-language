package main

import (
	"fmt"
	"strings"
)

func main() {
	// Initializing the Strings
	r := "Faizaaan"

	// Display the Strings
	fmt.Println("Given String: \n", r)

	// Using the Replace Function
	testresults := strings.Replace(r, "a", "", 3)

	// Display the ReplaceAll Output
	fmt.Println("\n After Replacing: \n", testresults)

}
