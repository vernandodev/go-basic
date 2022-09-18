package main

import "fmt"

// closure kemampuan sebuah function berinteraksi dengan data2 disekitarnya dalam scope yang sama *gunakan dengan bijak
func main() {
	name := "Andi"
	counter := 0

	increment := func ()  {
		name = "Budi"
		fmt.Println("Increment ")
		counter++
	}

	increment()
	increment()

	fmt.Println(name)
	fmt.Println(counter)
}