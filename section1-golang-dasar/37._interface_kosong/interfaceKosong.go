package main

import "fmt"

//interface kosong / {}

func Ups(i int) interface{} { // return interface{}
	if i == 1 {
		return 1
	}else if i == 2 {
		return true
	}else {
		return "ups"
	}
}

func main() {
	var data interface{} = Ups(1)
	fmt.Println(data)
}