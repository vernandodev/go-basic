package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

// function struct or method struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	eko := Customer {
		Name: "Eko",
		Address: "IDN",
		Age: 23,
	}

	// sayHello(eko, "joko")
	eko.sayHello("joko")

}