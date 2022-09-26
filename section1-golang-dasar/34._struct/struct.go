package main

import "fmt"

/**
Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan,
struct biasanya representasi data dalam program aplikasi yang dibuat,
data struct disimpan di dalam field,
sederhananya, struct adalah kumpulan field
*/

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	// cara 1
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "IDN"
	eko.Age = 30

	// cara 2
	// eko := Customer {
	// 	Name: "Eko",
	// 	Address: "IDN",
	// 	Age: 30,
	// }

	// cara 3
	// eko := Customer {"Eko", "IDN", 30}

	fmt.Println(eko)
	fmt.Println(eko.Address)
}