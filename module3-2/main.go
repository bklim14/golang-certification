package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animalMap := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hss"},
	}
	fmt.Println("Enter name and information separated by space")
	for {
		fmt.Print(">")
		scanner := bufio.NewReader(os.Stdin)
		line, _, _ := scanner.ReadLine()
		list := strings.Split(string(line), " ")
		if len(list) < 2 {
			fmt.Println("Wrong input")
			continue
		}
		chosenAnimal := animalMap[list[0]]
		if chosenAnimal.food == "" {
			fmt.Println("Animal does not exist")
		}
		switch list[1] {
		case "eat":
			chosenAnimal.Eat()
		case "move":
			chosenAnimal.Move()
		case "speak":
			chosenAnimal.Speak()
		default:
			fmt.Println("Info does not exist")
		}

	}

}
