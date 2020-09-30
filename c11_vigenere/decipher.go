package main

import (
	"fmt"
)

func main() {
	inputString := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	relativeIndex := func(indexValue int) int { return indexValue % 6 }
	for i := 0; i < len(inputString); i++ {
		char := inputString[i] - 65
		keyChar := keyword[relativeIndex(i)] - 65
		
		fmt.Printf(string((char - keyChar + 26) % 26 + 65))
	}
}