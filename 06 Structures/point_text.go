package sprint

import "fmt"

type Point struct {
	X, Y float64
	Text string
}

func PointText(p Point) Point {
	fText := fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return Point{X: p.X, Y: p.Y, Text: fText}
}
