package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	m := make(map[string]Animal)
	fmt.Println("Enter name and information separated by space")
	for {
		fmt.Print(">")
		scanner := bufio.NewReader(os.Stdin)
		line, _, _ := scanner.ReadLine()
		list := strings.Split(string(line), " ")
		if len(list) < 3 {
			fmt.Println("Missing inputs")
			continue
		}
		switch list[0] {
		case "newanimal":
			switch list[2] {
			case "cow":
				m[list[1]] = Cow{}
			case "bird":
				m[list[1]] = Bird{}
			case "snake":
				m[list[1]] = Snake{}
			default:
				fmt.Println("Animal not allowed")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			chosenAnimal := m[list[1]]
			if chosenAnimal == nil {
				fmt.Println("Animal does not exist")
				continue
			}
			switch list[2] {
			case "eat":
				chosenAnimal.Eat()
			case "move":
				chosenAnimal.Move()
			case "speak":
				chosenAnimal.Speak()
			default:
				fmt.Println("Info does not exist")
			}
		default:
			fmt.Println("Wrong command")
		}
	}
}
