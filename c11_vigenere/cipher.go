package main

import (
	"fmt"
)

func main() {
	inputString := "Hello, my name is Edd."
	keyword := "GOLANG"

	relativeIndex := func(indexValue int) int { return indexValue % len(keyword) }

	cipherFunc := func(isLower bool, index int, baseChar byte) {
		char := inputString[index] - baseChar
		var keyChar byte
		if isLower {
			keyChar = keyword[relativeIndex(index)] + 32 - baseChar	
		} else {
			keyChar = keyword[relativeIndex(index)] - baseChar
		}
		
		fmt.Printf(string((char + keyChar + 26) % 26 + baseChar))
	}

	for i := 0; i < len(inputString); i++ {
		if inputString[i] < 'z' && inputString[i] >= 'a' {
			cipherFunc(true, i, 97)
		} else if inputString[i] < 'Z' && inputString[i] >= 'A' {
			cipherFunc(false, i, 65)
		} else {
			fmt.Printf(string(inputString[i]))
		}
	}
}