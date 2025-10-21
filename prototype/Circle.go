package prototype

import "fmt"

// Circle 具体原型：圆形
type Circle struct {
	ID     string
	X, Y   int
	Radius int
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing Circle %s at (%d,%d) with radius %d\n",
		c.ID, c.X, c.Y, c.Radius)
}

func (c *Circle) GetID() string {
	return c.ID
}

func (c *Circle) Clone() Cloneable {
	// 创建新对象并复制所有字段
	return &Circle{
		ID:     c.ID + "_copy",
		X:      c.X,
		Y:      c.Y,
		Radius: c.Radius,
	}
}

/*
因为Go中默认的复制为浅拷贝，为了实现深拷贝，可以使用Json的序列化和反序列化实现
func (c *Circle) Clone() Cloneable{
	data,_:= json.Marshal(c)
	var newCircle Circle
	json.Unmarshal(data,&newCircle)
	newCircle.ID = c.ID + "_copy"
	return &newCircle
}
 */
