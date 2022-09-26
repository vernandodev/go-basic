package main

import "fmt"

// interface adalah tipe data abstract, tidak memilliki implementasi langsung (di interface tidak bisa bikin data di interface)
// di dalam interface hanya ada definisi2 method
type HasName interface {
	GetName() string // nama. tipe data harus sama dan parameter
}

func sayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

//struct
type Person struct {
	Name string
}

type Animal struct {
	Name string
}

//func struct
func (person Person) GetName() string { // disini
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	eko := Person {
		Name: "Eko",
	}

	kucing := Animal {
		Name: "Meong",
	}

	sayHello(eko)
	sayHello(kucing)
}