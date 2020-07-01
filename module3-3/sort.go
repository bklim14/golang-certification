package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sorter(arr []int, wg *sync.WaitGroup, c chan []int) {
	defer wg.Done()
	fmt.Printf("Goroutine - Unsorted Sorted Chunk %v\n", arr)
	sort.Ints(arr)
	c <- arr
}

func main() {
	fmt.Println("Enter minimum of 4 integers separated by a single space")
	fmt.Print(">")
	scanner := bufio.NewReader(os.Stdin)
	line, _, _ := scanner.ReadLine()
	list := strings.Split(string(line), " ")
	if len(list) < 4 {
		fmt.Println("Please input minimum of 4 integers")
		return
	}
	arr := make([]int, len(list))
	for index, val := range list {
		i, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Wrong input type")
			return
		}
		arr[index] = i
	}
	fmt.Printf("Unsorted Array %v\n", arr)
	middleIndex := len(list) / 2
	middleIndex2 := middleIndex / 2
	middleIndex3 := (middleIndex + len(list)) / 2

	c := make(chan []int, 4)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go sorter(arr[0:middleIndex2], &wg, c)
	wg.Add(1)
	go sorter(arr[middleIndex2:middleIndex], &wg, c)
	wg.Add(1)
	go sorter(arr[middleIndex:middleIndex3], &wg, c)
	wg.Add(1)
	go sorter(arr[middleIndex3:len(list)], &wg, c)

	var mainArr []int
	mainArr = append(mainArr, <- c...)
	mainArr = append(mainArr, <- c...)
	mainArr = append(mainArr, <- c...)
	mainArr = append(mainArr, <- c...)

	wg.Wait()
	sort.Ints(mainArr)
	fmt.Printf("Sorted Array %v\n", mainArr)
}