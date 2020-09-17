package main

type Object3d interface {
	Object2d

	MedianEdge() float64
}
