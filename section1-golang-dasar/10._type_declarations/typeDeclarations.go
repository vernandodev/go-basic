package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "232323232323232323"
	var marriedStatus Married = true
	fmt.Println(noKtpEko, marriedStatus)
}