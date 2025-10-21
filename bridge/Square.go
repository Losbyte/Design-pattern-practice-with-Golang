package bridge

import "fmt"

// Square 具体形状
type Square struct {
	*Shape
	side float64
}

func NewSquare(c Color, s float64) *Square {
	return &Square{
		Shape: NewShape(c),
		side:  s,
	}
}

func (s *Square) Draw() string {
	return fmt.Sprintf("Drawing Square with side %.2f in %s color",
		s.side, s.color.ApplyColor())
}
