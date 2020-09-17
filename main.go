package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Chair struct {
	x float64
	y float64
	z float64
}

func (chair Chair) canThrough(door Door) bool {
	return chair.MinEdge() <= door.MinEdge() && chair.MedianEdge() <= door.MaxEdge()
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
