// Go program to illustrate the
// concept of variadic function
package main

import (
	"fmt"
	"strings"
)

//Example 1:
// Variadic function to join strings
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {

	// zero argument
	fmt.Println(joinstr())

	// multiple arguments
	fmt.Println(joinstr("Manzoor", "Faisal"))
	fmt.Println(joinstr("Muhammad", "Manzoor", "Faisal"))
	fmt.Println(joinstr("M", "A", "N", "Z", "O", "O", "R"))

	// pass a slice in variadic function
	elements := []string{"Muhammad", "Manzoor", "Faisal"}
	fmt.Println(joinstr(elements...))
}
