package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}


func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) DistanceTo(other *Point) float64 {
	// расстояние - sqrt((x2 - x1)^2 + (y2 - y1)^2)
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(3, 4)
	p2 := NewPoint(7, 1)

	distance := p1.DistanceTo(p2)

	fmt.Printf("Distance between two points: %.2f\n", distance)
}
