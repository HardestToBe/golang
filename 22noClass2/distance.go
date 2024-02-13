package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}
type location struct {
	name      string
	lat, long float64
}
type distance struct {
	name string
	dist float64
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
func printLargestDist(dist []distance) {
	for _, value := range dist {
		fmt.Printf("%.2f km\n", value.dist)
	}
	var largestDist float64
	var largestDistName string
	for _, value := range dist {
		if largestDist < value.dist {
			largestDist = value.dist
			largestDistName = value.name
		}
	}
	fmt.Printf("Largest distance is from %v,and the number is %.2f km\n", largestDistName, largestDist)
}
func printDist(dist1 []distance, dist2 []distance) {
	fmt.Printf("%v:%.2f km\n", dist1[0].name, dist1[0].dist)
	fmt.Printf("%v:%.2f km\n", dist2[1].name, dist2[1].dist)
}

func main() {
	var mars = world{radius: 3389.5}
	var Earth = world{radius: 6371.0}
	locations := []location{
		{"spirit", -14.5684, 175.472636},
		{"opportunity", -1.9462, 354.4734},
		{"curiosity", -4.5895, 137.4417},
		{"insight", 4.5, 135.9},
		{"MtSharp", -5.08, 137.85},
		{"MtOlympus", 18.65, 226.2},
		{"London", 51.5000, -0.1333},
		{"Paris", 48.85, 2.35},
	}
	spiritToOpportunity := mars.distance(locations[0], locations[1])
	spiritToCuriosity := mars.distance(locations[0], locations[2])
	spiritToInsight := mars.distance(locations[0], locations[3])
	opportunityToOpportunity := mars.distance(locations[1], locations[2])
	opportunityToInsight := mars.distance(locations[1], locations[3])
	curiosityToInsight := mars.distance(locations[2], locations[3])
	MtSharpToMtOlympus := mars.distance(locations[4], locations[5])
	LondonToParis := Earth.distance(locations[6], locations[7])
	distMars := []distance{
		{"spiritToOpportunity", spiritToOpportunity},
		{"spiritToCuriosity", spiritToCuriosity},
		{"spiritToInsight", spiritToInsight},
		{"opportunityToOpportunity", opportunityToOpportunity},
		{"opportunityToInsight", opportunityToInsight},
		{"curiosityToInsight", curiosityToInsight},
		{"MtSharpToMtOlympus", MtSharpToMtOlympus},
	}
	distEarth := []distance{
		{"LondonToParis", LondonToParis},
	}
	printLargestDist(distMars)
	printDist(distEarth, distMars)
}
