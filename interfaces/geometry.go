package interfaces

import (
	"fmt"
	"math"
)

type Geometry interface {
	area()	float64
	perim()	float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) perim() float64 {
	return 2 * r.Width + 2 * r.Height
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.Radius
}

func Measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("Area =", g.area())
	fmt.Println("Parameter =", g.perim())
}