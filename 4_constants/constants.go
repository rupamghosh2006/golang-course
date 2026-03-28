package main

import "fmt"

//const age = 30
// we can declare variable and constants outside main func but we cant declare short hand

func main(){
	//const name string = "golang"

	//const age = 30
	
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}