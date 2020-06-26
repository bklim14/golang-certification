package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(seq []int, index int) {
	seq[index+1], seq[index] = seq[index], seq[index+1]
}

func BubbleSort(seq []int) {
	for i := 0; i < len(seq); i++ {
		for j := 0; j < len(seq)-i-1; j++ {
			if seq[j] > seq[j+1] {
				Swap(seq, j)
			}
		}
	}
}

func main() {
	fmt.Println("Enter sequence of up to 10 integers separated by a single space:")
	scanner := bufio.NewReader(os.Stdin)
	line, _, _ := scanner.ReadLine()
	numberList := strings.Split(strings.Trim(string(line), " "), " ")
	if len(numberList) > 10 {
		fmt.Println("Sequence greater than 10")
		return
	}
	seq := make([]int, len(numberList))
	for index, val := range numberList {
		number, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Invalid number in sequence")
			return
		}
		seq[index] = number
	}
	BubbleSort(seq)
	fmt.Println(seq)
}
