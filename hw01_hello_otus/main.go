package main

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
)

func main() {
	name := "OTUS"
	helloString := fmt.Sprintf("Hello, %s!", name)

	fmt.Println(reverse.String(helloString))
}
