package main

import (
	"math"
)

type Door struct {
	Object2d

	x float64
	y float64
}

func (door Door) MinEdge() float64 {
	return math.Min(door.x, door.y)
}

func (door Door) MaxEdge() float64 {
	return math.Max(door.x, door.y)
}
