package mediator

import "fmt"

// ControlTower 具体中介者（塔台）
type ControlTower struct {
	isRunwayAvailable bool
	aircrafts         []Aircraft
}

func NewControlTower() *ControlTower {
	return &ControlTower{
		isRunwayAvailable: true,
	}
}

func (c *ControlTower) Register(aircraft Aircraft) {
	c.aircrafts = append(c.aircrafts, aircraft)
}

func (c *ControlTower) Notify(sender interface{}, event string) {
	switch event {
	case "landingRequest":
		if c.isRunwayAvailable {
			c.isRunwayAvailable = false
			sender.(Aircraft).NotifyLandingApproved()
		} else {
			fmt.Printf("塔台回复 %s: 跑道占用，请等待\n", sender.(Aircraft).Name())
		}
	case "takeoffRequest":
		if c.isRunwayAvailable {
			c.isRunwayAvailable = false
			sender.(Aircraft).NotifyTakeoffApproved()
		} else {
			fmt.Printf("塔台回复 %s: 跑道占用，请等待\n", sender.(Aircraft).Name())
		}
	case "landingCompleted":
		fmt.Printf("塔台记录: %s 已降落\n", sender.(Aircraft).Name())
		c.isRunwayAvailable = true
	case "takeoffCompleted":
		fmt.Printf("塔台记录: %s 已起飞\n", sender.(Aircraft).Name())
		c.isRunwayAvailable = true
	}
}
