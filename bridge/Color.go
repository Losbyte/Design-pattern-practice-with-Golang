package bridge

// Color 颜色接口 - 实现部分
type Color interface {
	ApplyColor() string
}

// Red 具体颜色实现
type Red struct{}

func (r *Red) ApplyColor() string {
	return "Red"
}

// Green 具体颜色实现
type Green struct{}

func (g *Green) ApplyColor() string {
	return "Green"
}

// Blue 具体颜色实现
type Blue struct{}

func (b *Blue) ApplyColor() string {
	return "Blue"
}
