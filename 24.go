package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) DotDistance(distance Point) float64 {
	// просто по математической формуле и добавил abs, потому что дистанция не может быть отрицательной и корень тоже в данном случае
	return math.Sqrt(math.Abs(math.Pow(float64(p.x-distance.x), 2) + math.Pow(float64(p.y-distance.y), 2)))
}

func task24() {
	a := NewPoint(0, 1)
	b := NewPoint(2, -2)

	// от точки считаю дистанцию до точки b
	dist := a.DotDistance(*b)

	fmt.Println(dist)
}
