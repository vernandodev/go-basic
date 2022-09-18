package main

import "fmt"

// fungsi yang memanggil function dirinya sendiri

// contoh yang belum memakai recursive function
func factorialLoop(value int) int {
	result := 1
	for i:= value; i > 0; i-- {
		result *= i
	}
	return result
}

// recursive funtion *harus ada moment berhenti
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	}else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	loopRecursive := factorialRecursive(5)
	fmt.Println(loopRecursive)
}