package bridge

// Shape 形状抽象类 - 抽象部分
type Shape struct {
	color Color
}

// NewShape 构造函数
func NewShape(c Color) *Shape {
	return &Shape{color: c}
}

// Draw 绘制方法，由具体形状实现
func (s *Shape) Draw() string {
	return ""
}
