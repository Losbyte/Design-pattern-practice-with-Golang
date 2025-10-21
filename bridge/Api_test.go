package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T){
	// 创建颜色实例
	red := &Red{}
	green := &Green{}
	blue := &Blue{}

	// 创建形状实例
	circle := NewCircle(red, 5.0)
	square := NewSquare(green, 10.0)
	bigCircle := NewCircle(blue, 15.0)

	// 绘制图形
	fmt.Println(circle.Draw())
	fmt.Println(square.Draw())
	fmt.Println(bigCircle.Draw())
}