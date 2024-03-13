package main

import (
	"fmt"
)

func main() {
	var passwordLength int
	fmt.Println("Enter password length: ")
	fmt.Scan(&passwordLength)
	numberPassword := numberGeneration(passwordLength)
	stringPassword := stringGenerator(passwordLength)
	fmt.Println(numberPassword)
	fmt.Println(stringPassword)
}
