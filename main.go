package main

import (
	"flag"
	"fmt"
)

const procedureName = "com.things.echo"
const url = "ws://0.0.0.0:8080/ws"

func main() {
	fmt.Print("Put 0 for caller and 1 for callee ")
	urlString := flag.String("url", url, "url to connect")
	procedureString := flag.String("proc", procedureName, "procedure name")
	flag.Parse()
	var input int
	fmt.Scanln(&input)
	if input == 0 {
		caller(*urlString, *procedureString)
	} else if input == 1 {
		callee(*urlString, *procedureString)
	}
}
