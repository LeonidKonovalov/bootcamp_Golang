package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func geometryArea(g geometry) float64 {
	return g.area()
}

func main() {
	fmt.Println(geometryArea(&rect{width: 3, height: 4}))
	fmt.Println(geometryArea(&circle{radius: 5}))
}
