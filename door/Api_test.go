package door

import (
	"fmt"
	"testing"
)

func TestMovie(t *testing.T){
	homeTheater := NewHomeTheaterFacade()

	// 使用外观简化操作
	fmt.Println("=== 使用外观模式 ===")
	homeTheater.WatchMovie("The Matrix")
	homeTheater.EndMovie()

	// 对比：不使用外观模式的复杂操作
	fmt.Println("=== 不使用外观模式（复杂操作）===")
	amp := &Amplifier{}
	projector := &Projector{}
	dvd := &DVDPlayer{}
	lights := &Lights{}
	screen := &Screen{}

	// 需要手动操作每个设备
	lights.Dim(10)
	screen.Down()
	projector.On()
	projector.SetInput("DVD")
	amp.On()
	amp.SetVolume(5)
	dvd.On()
	dvd.Play("Inception")
	fmt.Println("Movie is ready!")

	// 对比显示外观模式的价值
	fmt.Println("=== 外观模式的优势 ===")
	fmt.Println("1. 简化客户端代码")
	fmt.Println("2. 隐藏子系统复杂性")
	fmt.Println("3. 提供统一接口")
	fmt.Println("4. 降低耦合度")
}

