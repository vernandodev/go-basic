package main

import "fmt"

func main() {
	name := "joko"

	// If short statement
	if length := len(name); length>5 {
		fmt.Println("Terlalu panjang")
	}else {
		fmt.Println("Nama sudah benar")
	}

	if name == "eko"{
		fmt.Println("Benar")
	}

	if name == "eko"{
		fmt.Println("Benar")
	}else {
		fmt.Println("Salah")
	}

	if name == "eko" {
		fmt.Println("Hello", name)
	}else if name == "joko" {
		fmt.Println("Hello joko")
	}else {
		fmt.Println("Salah")
	}



}