package main

import "fmt"

func getHello() (string, string, string){
	return "Eko", "Khannedy", "test"
}

func main() {
	firstname, _, test := getHello()
	fmt.Println(firstname, test)
}