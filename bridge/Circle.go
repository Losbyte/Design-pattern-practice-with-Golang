package bridge

import "fmt"

// Circle 具体形状
type Circle struct {
	*Shape
	radius float64
}

func NewCircle(c Color, r float64) *Circle {
	return &Circle{
		Shape:  NewShape(c),
		radius: r,
	}
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Drawing Circle with radius %.2f in %s color",
		c.radius, c.color.ApplyColor())
}
