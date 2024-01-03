package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(hello("Stefan"))
}

func hello(name string) string {
	message := fmt.Sprintf("Hello, %v.", name)
	return message
}
