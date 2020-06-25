package main

import (
	"fmt"
	"sort"
)

func main() {
	var input, sliceLen, index int
	sliceLen = 3
	s := make([]int, sliceLen)
	for {
		fmt.Print("Enter integer:")
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		if index >= sliceLen {
			s = append(s, input)
		} else {
			s[index] = input
		}
		sortedS := make([]int, len(s))
		copy(sortedS, s)
		sort.Ints(sortedS)
		fmt.Println(sortedS)
		index++
	}
}
