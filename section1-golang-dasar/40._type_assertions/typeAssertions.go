package main

import "fmt"

// type assertion digunakan untuk merubah tipe data
func random() interface{} {
	return "Ups"
}

func main() {
	// var result interface{} = random()
	// var resultString string = result.(string)

	// fmt.Println(resultString)

	// type assertion menggunakan switch

	var result interface{} = random()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}


