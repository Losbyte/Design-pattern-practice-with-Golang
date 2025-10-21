package prototype

import "fmt"

// Rectangle 具体原型：矩形
type Rectangle struct {
	ID          string
	X, Y        int
	Width, Height int
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing Rectangle %s at (%d,%d) with width %d and height %d\n",
		r.ID, r.X, r.Y, r.Width, r.Height)
}

func (r *Rectangle) GetID() string {
	return r.ID
}

func (r *Rectangle) Clone() Cloneable {
	// 创建新对象并复制所有字段
	return &Rectangle{
		ID:     r.ID + "_copy",
		X:      r.X,
		Y:      r.Y,
		Width:  r.Width,
		Height: r.Height,
	}
}
