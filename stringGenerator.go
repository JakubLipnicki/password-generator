package main

import (
	"fmt"
	"math/rand/v2"
)

func stringGenerator(passwordLength int) (password string) {
	characters := "vnaefghijEFGopqwxyzABIJKCDLMbcdNHklmrstuOPQRSTUVWXYZ"
	for i := 0; i < passwordLength; i++ {
		password += fmt.Sprint(string(characters[rand.IntN(len(characters))]))
	}
	return password
}
