package main

import (
	"database/database.go"
	"fmt"
)

func p() {
	result := database.GetDatabase()
	fmt.Println(result)
}