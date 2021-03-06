package main

import "math"

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y2)

	return l * w
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r // PI * r^2
}

type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{0, 0, 5}

}
