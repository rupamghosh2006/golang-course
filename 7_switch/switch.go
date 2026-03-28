package main

import (
	"fmt"
	"time"
)

func main() {

	//simple switch
	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")	
	default:
		fmt.Println("Other")
	}

	// we dont have to use break statement in GO

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}


	// type switch
	whoAmI := func(i interface{}){
		switch t:= i.(type) {
		case int:
			fmt.Println("It is an Integer")
		case string:
			fmt.Println("It is a strin")
		case bool:
			fmt.Println("It is a boolean")
		default:
			fmt.Println("Other", t)	
		}
	}

	whoAmI(50)
}