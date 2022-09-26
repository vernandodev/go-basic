package main

import "fmt"

// secara default di golang semua variable itu di passing by value, bukan by reference
// artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya

type Address struct {
	City, Province, Country string
	Age int
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := &address1

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia", 87}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	address2 = &Address{"Subang", "Jawa Barat", "Indonesia", 98}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}