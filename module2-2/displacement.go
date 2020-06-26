package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return ((a * (t * t) / 2) + (v * t) + s)
	}
}
func main() {
	var a, v, s, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v, s)
	fmt.Printf("Displacement after %f seconds = %f\n", t, fn(t))
}
