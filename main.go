package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Object interface {
}

type Object2d interface {
	Object

	MinEdge() float64
	MaxEdge() float64
}

type Object3d interface {
	Object2d

	MedianEdge() float64
}

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

func main() {
	char := Chair{x: 1, y: 20, z: 15}
	door := Door{x: 10, y: 1}

	if char.canThrough(door) {
		fmt.Println("通れます")
		os.Exit(0)
	}

	fmt.Println("通れません")
	os.Exit(1)
}
