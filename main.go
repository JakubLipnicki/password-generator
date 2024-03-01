package main

import (
	"fmt"
)

func main() {
	var passwordLength int
	fmt.Println("Enter password length: ")
	fmt.Scan(&passwordLength)
	var password string
	password = numberGeneration(passwordLength)
	fmt.Println(password)
}
