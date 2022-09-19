package main

import "fmt"

func main() {
	var name1 = "andi"
	var name2 = "richo"
	var perbandinganNama bool = name1 == name2
	fmt.Println(perbandinganNama)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value2 > value1)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}