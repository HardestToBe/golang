package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}
type location struct {
	lat, long float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
func printLargestDist(dist []float64) {
	for _, value := range dist {
		fmt.Printf("%.2f km\n", value)
	}
	var largestDist float64
	for _, value := range dist {
		if largestDist < value {
			largestDist = value
		}
	}
	fmt.Printf("largest distance:%.2f km\n", largestDist)
}
func printDist(dist1 []float64, dist2 []float64) {
	fmt.Printf("Distance of London To Paris:%.2f km\n", dist1[0])
	fmt.Printf("Distance of MtSharp To MtOlympus:%.2f km\n", dist2[6])
}

func main() {
	var mars = world{radius: 3389.5}
	var Earth = world{radius: 6371.0}
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	curiosity := location{-4.5895, 137.4417}
	insight := location{4.5, 135.9}
	MtSharp := location{-5.08, 137.85}
	MtOlympus := location{18.65, 226.2}
	distMars := []float64{
		mars.distance(spirit, opportunity),
		mars.distance(spirit, curiosity),
		mars.distance(spirit, insight),
		mars.distance(opportunity, curiosity),
		mars.distance(opportunity, insight),
		mars.distance(curiosity, insight),
		mars.distance(MtSharp, MtOlympus),
	}
	London := location{51.5000, -0.1333}
	Paris := location{48.85, 2.35}

	disEarth := []float64{
		Earth.distance(London, Paris),
	}
	printLargestDist(distMars)
	printDist(disEarth, distMars)
}
