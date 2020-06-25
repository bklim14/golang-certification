package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name:")
	bName, _, _ := scanner.ReadLine()
	m["name"] = string(bName)
	fmt.Print("Enter address:")
	bAddress, _, _ := scanner.ReadLine()
	m["address"] = string(bAddress)
	b, _ := json.Marshal(m)
	fmt.Println(string(b))
}