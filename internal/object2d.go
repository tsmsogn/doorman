package internal

type Object2d interface {
	Object

	MinEdge() float64
	MaxEdge() float64
}
