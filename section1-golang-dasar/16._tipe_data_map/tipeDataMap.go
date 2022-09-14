package main

import "fmt"

func main() {
	person := map[string]string{
		"Name" : "Richo",
		"Address" : "Indo",
	}



	fmt.Println(person)
	fmt.Println(person["Name"])
	fmt.Println(len(person)) // len(map)
	fmt.Println(person["Address"]) //map[key]
	person["Name"] = "Programmer" // map[key] = value
	fmt.Println(person["Name"])

	book := map[string]string{
		"Title" : "Ayahku seorang pelaut",
		"Writer" : "John Steel",
	}
	book["Years"] = "2019"
	fmt.Println(book)

	delete(book, "Years")
	fmt.Println(book)
}