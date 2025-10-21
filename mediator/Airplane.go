package mediator

import "fmt"

// Airplane 具体飞机组件
type Airplane struct {
	name     string
	mediator Mediator
}

func NewAirplane(name string, mediator Mediator) *Airplane {
	return &Airplane{
		name:     name,
		mediator: mediator,
	}
}

func (a *Airplane) Name() string {
	return a.name
}

func (a *Airplane) RequestLanding() {
	fmt.Printf("%s 请求降落\n", a.name)
	a.mediator.Notify(a, "landingRequest")
}

func (a *Airplane) RequestTakeoff() {
	fmt.Printf("%s 请求起飞\n", a.name)
	a.mediator.Notify(a, "takeoffRequest")
}

func (a *Airplane) NotifyLandingApproved() {
	fmt.Printf("%s 收到降落许可\n", a.name)
}

func (a *Airplane) NotifyTakeoffApproved() {
	fmt.Printf("%s 收到起飞许可\n", a.name)
}
