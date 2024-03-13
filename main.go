package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var passwordLength, charactersType int
	password := ""
	fmt.Println("Enter password length: ")
	fmt.Scan(&passwordLength)
	fmt.Println("What characters would you like to include in your password")
	fmt.Println(`	Numbers: type 1
	Letters: type 2
	Symbols: type 3
	Combination: type numbers in ascending order e.g. (13)`)
	fmt.Scan(&charactersType)
	if charactersType == 1 {
		fmt.Println(numberGeneration(passwordLength))
		return
	}
	if charactersType == 2 {
		fmt.Println(stringGenerator(passwordLength))
		return
	}
	if charactersType == 3 {
		fmt.Println(symbolGenerator(passwordLength))
		return
	}
	typeToString := fmt.Sprint(charactersType)
	for i := 0; i < passwordLength; i++ {
		option := string(typeToString[rand.IntN(len(typeToString))])
		switch option {
		case "1":
			password += numberGeneration(1)
		case "2":
			password += stringGenerator(1)
		case "3":
			password += symbolGenerator(1)
		default:
			password += "test"
		}
	}
	fmt.Println(password)

	// numberPassword := numberGeneration(passwordLength)
	// stringPassword := stringGenerator(passwordLength)
	// specialPassword := specialCharacterGenerator(passwordLength)
	// fmt.Println(numberPassword)
	// fmt.Println(stringPassword)
	// fmt.Println(specialPassword)
}
