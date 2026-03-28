package main

import "fmt"

func main() {
	age := 18

	// if age >= 18 {
	// 	fmt.Println("Person is an adult")
	// } else {
	// 	fmt.Println("Person is not an adult")
	// }

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12{
		fmt.Println("Person is teenager")
	} else {
		fmt.Println("Person is a kid")
	}

	//go does not have ternary operator
}