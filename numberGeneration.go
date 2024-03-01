package main

import (
	"fmt"
	"math/rand/v2"
)

func numberGeneration(passwordLength int) (password string) {
	for i := 1; i <= passwordLength; i++ {
		password += fmt.Sprint(rand.IntN(10))
	}
	return password
}
