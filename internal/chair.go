package internal

import (
	"math"
	"sort"
)

type Chair struct {
	Object3d

	X float64
	Y float64
	Z float64
}

func (chair Chair) CanThrough(object2d Object2d) bool {
	return chair.MinEdge() <= object2d.MinEdge() && chair.MedianEdge() <= object2d.MaxEdge()
}

func (chair Chair) MedianEdge() float64 {
	edges := []float64{chair.X, chair.Y, chair.Z}
	sort.Float64s(edges)

	return edges[1]
}

func (chair Chair) MinEdge() float64 {
	return math.Min(chair.X, chair.Y)
}
