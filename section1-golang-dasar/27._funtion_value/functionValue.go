package main

import "fmt"

// funtion disimpan dalam variabel

func getGoodBye(name string) string {
	return "Hello " + name
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("eko"))
}