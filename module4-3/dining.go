package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number int
	left, right *Chopstick
}

func (p Philosopher) Eat(c chan int) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		<- c
		p.left.Lock()
		p.right.Lock()
		fmt.Printf("starting to eat %d\n", p.number)
		p.left.Unlock()
		p.right.Unlock()
		fmt.Printf("finishing eating %d\n", p.number)
	}
}

var wg = sync.WaitGroup{}

func invite(c chan int) {
	defer wg.Done()
	for i := 0; i < 15; i++{
		c <- i
	}
}

func main() {
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			i + 1,
			chopsticks[i],
			chopsticks[(i+1) % 5],
		}
	}

	c := make(chan int, 2)
	wg.Add(1)
	go invite(c)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].Eat(c)
	}
	wg.Wait()
}
