package mediator

import (
	"fmt"
	"testing"
	"time"
)

func TestMediator(t *testing.T) {
	// 创建中介者（塔台）
	tower := NewControlTower()

	// 创建飞机并注册到塔台
	flight1 := NewAirplane("航班CA123", tower)
	flight2 := NewAirplane("航班UA456", tower)
	flight3 := NewAirplane("航班BA789", tower)

	tower.Register(flight1)
	tower.Register(flight2)
	tower.Register(flight3)

	// 模拟飞机请求
	fmt.Println("=== 机场交通控制模拟 ===")

	flight1.RequestLanding()
	flight2.RequestLanding() // 应该被拒绝

	time.Sleep(2 * time.Second) // 模拟降落过程
	tower.Notify(flight1, "landingCompleted")

	flight2.RequestLanding() // 现在应该被允许
	flight3.RequestTakeoff() // 应该被拒绝

	time.Sleep(2 * time.Second) // 模拟降落过程
	tower.Notify(flight2, "landingCompleted")

	flight3.RequestTakeoff() // 现在应该被允许

	time.Sleep(1 * time.Second) // 模拟起飞过程
	tower.Notify(flight3, "takeoffCompleted")

	fmt.Println("=== 中介者模式解决的问题 ===")
	fmt.Println("1. 飞机间无需直接通信，降低耦合")
	fmt.Println("2. 塔台集中管理所有交互，避免混乱")
	fmt.Println("3. 跑道状态由塔台统一维护")
	fmt.Println("4. 新增飞机不影响现有逻辑")
}
