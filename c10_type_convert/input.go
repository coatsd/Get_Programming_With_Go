package main

import (
	"fmt"
)

func main() {
	launchString := "no"
	var isTrue bool
	var isError bool

	switch launchString {
	case "no":
		isTrue = false
	case "n":
		isTrue = false
	case "false":
		isTrue = false
	case "f":
		isTrue = false
	case "0":
		isTrue = false
	case "yes":
		isTrue = true
	case "y":
		isTrue = true
	case "true":
		isTrue = true
	case "t":
		isTrue = true
	case "1":
		isTrue = true
	default:
		isError = true
	}

	if isError {
		fmt.Printf("An error has occurred. The string %v cannot be converted to a bool.", launchString)
	} else {
		resultString := func() string {
			if isTrue {
				return "true"
			}
			return "false"
		}
		fmt.Printf("launchString = %v , result = %v", launchString, resultString())
	}
}