package main

import "fmt"

type geo interface {
	area() float64
	persim() float64
}

type react struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r react) area() float64 {
	return r.width * r.height
}
func (r react) persim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return 2 * 3.14 * c.radius
}
func (c circle) persim() float64 {
	return 3.14 * c.radius * c.radius
}

func measure(g geo) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.persim())
}

func main() {
	r := react{width: 5, height: 10}
	c := circle{radius: 7}

	measure(r)
	measure(c)
}
