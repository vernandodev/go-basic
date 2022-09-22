package main

import "fmt"

// nil hanya bisa digunakan di beberapa tipe data, interface, function, map, slice, pointer, channel

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}else {
		return map[string]string {
			"name" : name,
		}
	}
}

func main() {
	var person map[string]string = newMap("Eko")

	if person == nil {
		fmt.Println("Data kosong")
	}else {
		fmt.Println(person)
	}

	// var person map[string] string = nil
	// fmt.Println(person)

	// person := newMap("")
	// fmt.Println(person)
}