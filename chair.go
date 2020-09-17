package main

import (
	"math"
	"sort"
)

type Chair struct {
	Object3d

	x float64
	y float64
	z float64
}

func (chair Chair) canThrough(object2d Object2d) bool {
	return chair.MinEdge() <= object2d.MinEdge() && chair.MedianEdge() <= object2d.MaxEdge()
}

func (chair Chair) MedianEdge() float64 {
	edges := []float64{chair.x, chair.y, chair.z}
	sort.Float64s(edges)

	return edges[1]
}

func (chair Chair) MinEdge() float64 {
	return math.Min(chair.x, chair.y)
}
