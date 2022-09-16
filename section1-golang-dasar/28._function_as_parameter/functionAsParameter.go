package main

import (
	"fmt"
)

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string{
	if name == "Kasar" {
		return "..."
	}else {
		return name
	}
}

// menggunakan type declaration untuk memperpendek penggunaan

type Filter func(string)string

func inputWordWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func main() {
	// penerapan
	//cara 1
	sayHelloWithFilter("Eko", spamFilter)
	inputWordWithFilter("Richo", spamFilter)

	//cara2
	filter := spamFilter
	sayHelloWithFilter("Kasar", filter)
}