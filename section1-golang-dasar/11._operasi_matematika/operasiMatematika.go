package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var positive = +100
	var negative = -100
	fmt.Println(positive, negative)

}