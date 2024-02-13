package main

import (
	"fmt"
	"math"
)

type gps struct {
	startLocation location
	endLocation   location
	world         world
}
type location struct {
	name      string
	lat, long float64
}
type world struct {
	radius float64
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
func (l location) description() string {
	return fmt.Sprintf("%v (latitude:%.2f,longtitude:%.2f)", l.name, l.lat, l.long)
}
func (w world) distance(start, end location) float64 {
	s1, c1 := math.Sincos(rad(start.lat))
	s2, c2 := math.Sincos(rad(end.lat))
	clong := math.Cos(rad(start.long - end.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}
func (g gps) distance() float64 {
	return g.world.distance(g.startLocation, g.endLocation)
}
func (g gps) message() string {
	return fmt.Sprintf("%.2f km to %v", g.distance(), g.endLocation.description())
}

type rover struct {
	gps
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{"布莱德柏利着陆地", -4.5895, 137.4417}
	elysium := location{"埃律西昂平原", 4.5, 135.9}
	gps := gps{
		world:         mars,
		startLocation: bradbury,
		endLocation:   elysium,
	}
	curiosity := rover{
		gps: gps,
	}
	fmt.Println(curiosity.message())

}
