package main

import "fmt"

func main() {
	sayHelloTo("Eko", "kURNIAWAN")
}

func sayHelloTo(firstName string, lastname string) {
	fmt.Println("Hello", firstName, lastname)
}