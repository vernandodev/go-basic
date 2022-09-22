package main

import "fmt"

// type assertion digunakan untuk merubah tipe data
func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)

	fmt.Println(resultString)

	// type assertion menggunakan switch

}


