package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)
type name struct {
	firstName string
	lastName string
}

func main(){
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file name:")
	bLine, _, _ := scanner.ReadLine()
	file, err := os.Open(string(bLine))
	if err != nil {
		fmt.Println("File does not exist")
	}
	scanner = bufio.NewReader(file)
	var nameList []name
	for {
		line, _, err := scanner.ReadLine()
		if err != nil {
			break
		}
		s := bytes.Split(line, []byte(" "))
		firstNameBytes := make([]byte, 20)
		copy(firstNameBytes, s[0])
		lastNameBytes := make([]byte, 20)
		copy(lastNameBytes, s[1])
		nameList = append(nameList, name{string(firstNameBytes),string(lastNameBytes)})
	}

	for _, val := range nameList {
		fmt.Println(val.firstName, val.lastName)
	}
}