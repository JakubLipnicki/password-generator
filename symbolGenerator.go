package main

import (
	"fmt"
	"math/rand/v2"
)

func symbolGenerator(passwordLength int) (password string) {
	characters := "~`!@#$%^&*()_-+={[}]|:<,>.?/;\"'\\"

	for i := 0; i < passwordLength; i++ {
		password += fmt.Sprint(string(characters[rand.IntN(len(characters))]))
	}

	return password
}
