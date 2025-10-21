package prototype

import (
	"fmt"
	"testing"
)

func TestProtype(t *testing.T){
	// 初始化原型管理器
	manager := NewGraphicManager()

	// 注册原型
	manager.Register("default_circle", &Circle{
		ID:     "circle_1",
		X:      10,
		Y:      10,
		Radius: 5,
	})

	manager.Register("default_rect", &Rectangle{
		ID:     "rect_1",
		X:      20,
		Y:      20,
		Width:  8,
		Height: 6,
	})

	// 克隆图形
	circle := manager.CreateGraphic("default_circle")
	rect := manager.CreateGraphic("default_rect")

	// 使用克隆的图形
	circle.Draw()
	rect.Draw()

	// 验证克隆结果
	fmt.Printf("Original Circle ID: %s\n", manager.prototypes["default_circle"].GetID())
	fmt.Printf("Cloned Circle ID: %s\n", circle.GetID())
}


