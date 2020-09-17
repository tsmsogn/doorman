package internal

import (
	"math"
)

type Door struct {
	Object2d

	X float64
	Y float64
}

func (door Door) MinEdge() float64 {
	return math.Min(door.X, door.Y)
}

func (door Door) MaxEdge() float64 {
	return math.Max(door.X, door.Y)
}
