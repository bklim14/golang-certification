package main

import "fmt"

func main() {
    var input float64
    fmt.Print("Enter floating point number: ")
    _, err := fmt.Scan(&input)
    if err != nil {
        fmt.Println("Wrong input type.")
        return
    }
    var trunc int = int(input)
    fmt.Printf("Original Input - %f, Truncated Result = %d\n", input, trunc)
}
