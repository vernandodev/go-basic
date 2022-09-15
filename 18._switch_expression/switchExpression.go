package main

import "fmt"

func main() {
	name := "eko"

	switch name {
	case "eko" :
		fmt.Println("Benar")
	case "joko" :
		fmt.Println("Salah")
	case "trio" :
		fmt.Println("Salah")
	default :
		fmt.Println("Default")
	}

	//short switch statement
	switch length := len(name); length>5 {
	case true:
		fmt.Printf("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//switch tanpa kondisi
	length := len(name)
	switch{
	case length>5:
		fmt.Println("Nama terlalu panjang")
	case length<5:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Default")
	}
}