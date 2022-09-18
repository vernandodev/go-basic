package main

import "fmt"

// tidak perlu membuat func tapi langsung dibuat di variable dan parameter

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked " + name)
	}else {
		fmt.Println("Welcome " + name)
	}
}

// func blaklistAdmin (name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot (name string) bool {
// 	return name == "root"
// }

func main() {
	blacklist := func (name string) bool  {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("root", blacklist)

	// atau bisa langsung ditempatkan anonymous functionya didalam parameter

	registerUser("admin", func (name string) bool  {
		return name == "admin"
	})
}