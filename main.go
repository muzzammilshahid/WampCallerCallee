package main

import "fmt"

const procedureName = "com.things.echo"

func main() {
	fmt.Print("Put 0 for caller and 1 for callee ")
	var input int
	fmt.Scanln(&input)
	if input == 0 {
		caller()
	} else if input == 1 {
		callee()
	}
}
