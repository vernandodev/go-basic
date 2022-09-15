package main

import "fmt"

func main() {

	// for with statement
	for counter := 3; counter <= 10; counter++{
		fmt.Println("Perulangan", counter)
	}

	slice := []string{"Eko", "Kurniawan", "test", "budi"}
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}

	// for Range
	for index, value := range slice{
		fmt.Println("Index", index, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "richo"
	person["addres"] = "indo"

	for key, value := range person{
		fmt.Println(key, "=", value)
	}

	// test
	n := 1
	for n<5 {
		n *= 2 // 1,2,4
	}
	fmt.Println(n) // 8 (1*2*2*2)

	m:= 1
	for m<10 {
		m = m * 5
	}
	fmt.Println(m)

	o := 1
	for o < 100 {
		o += o
	}
	fmt.Println(o)



}