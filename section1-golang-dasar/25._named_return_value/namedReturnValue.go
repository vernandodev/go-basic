package main

import "fmt"

func getCompleteName() (firstname int, middlename int, lastname string) {
	firstname = 12345454535352323
	middlename = 123
	lastname = "richo"

	return
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}