package main

import (
	"fmt"
	"math/rand"
	"time"
)

type animal interface {
	move() string
	eat() string
}
type xenon struct {
	name string
}

func (x xenon) String() string {
	return x.name
}
func (x xenon) move() string {
	switch rand.Intn(2) {
	case 0:
		return "fuck you organism!"
	default:
		return "zizizi"
	}
}
func (x xenon) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "organism"
	default:
		return "silicon"
	}
}

type boron struct {
	name string
}

func (b boron) String() string {
	return b.name
}
func (b boron) move() string {
	switch rand.Intn(2) {
	case 0:
		return "putaputa"
	default:
		return "shishi"
	}
}
func (b boron) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "water"
	default:
		return "bofu"
	}
}
func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v is having %v.\n", a, a.eat())
	case 1:
		fmt.Printf("%v is moving and the sound is %v.\n", a, a.move())
	}

}
func main() {
	dayHour := 24
	days := 3
	sunRise := 8
	sunFall := 18
	animal := []animal{
		xenon{name: "Xenon"},
		boron{name: "Boron"},
	}
	for d := 0; d < days; d++ {
		for h := 0; h < dayHour; h++ {
			if h >= sunRise && h <= sunFall {
				i := rand.Intn(len(animal))
				step(animal[i])
			} else {
				fmt.Println("Animals are sleeping.")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
