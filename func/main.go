package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type rect struct {
	with , height float64
}

type circle struct {
	radius float64
}

func (r *rect)area() float64 {
	return r.with * r.height
}

func (c * circle)area() float64 {
	return math.Pi * c.radius * c.radius
}

func main(){
	rec := new(rect)
	rec.with = 50
	rec.height = 18

	c := new(circle)
	c.radius = 10

	fmt.Println(rec.area())
	fmt.Println(c.area())
}
