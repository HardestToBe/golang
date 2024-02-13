package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}
type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}
func main() {
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})
	London := newLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	Paris := newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	MtSharp := newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	MtOlympus := newLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})
	fmt.Printf("%.4f\n", spirit)
	fmt.Printf("%.4f\n", opportunity)
	fmt.Println(curiosity)
	fmt.Println(insight)
	fmt.Printf("%.4f\n", London)
	fmt.Println(Paris)
	fmt.Println(MtSharp)
	fmt.Println(MtOlympus)
}
