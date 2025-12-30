package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello go")

	var name string
	name = "go lang"
	sayHello(name)

	var number int
	number = 2

	var secondNumber = 5

	result := 7

	print(result, " - ", number, "-", secondNumber)

}

func sayHello(name string) {

	fmt.Println("hello" + name)

}
