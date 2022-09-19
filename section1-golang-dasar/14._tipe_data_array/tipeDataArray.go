package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Richo"
	names[1] = "Andi"
	names[2] = "Budi"
	fmt.Println(names[0])
	fmt.Println(names[2])

	var country = [3]string{
		"indonesia",
		"thailand",
		"malaysia",
	}
	fmt.Println(country[2])

	var values = [3]int{
		90,
		80,
		90,
	}
	fmt.Println(values[0])

	fmt.Println(len(values), len(names))



}