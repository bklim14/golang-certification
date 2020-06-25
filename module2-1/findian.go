package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Enter string: ")
	fmt.Scan(&input)

	inputUpper := strings.ToUpper(input)
	if strings.HasPrefix(inputUpper, "I") &&
		strings.HasSuffix(inputUpper, "N") &&
		strings.Contains(inputUpper, "A"){
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
