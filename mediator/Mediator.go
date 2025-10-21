package mediator

// Mediator 中介者接口
type Mediator interface {
	Notify(sender interface{}, event string)
}

// Aircraft 飞机组件接口
type Aircraft interface {
	Name() string
	RequestLanding()
	RequestTakeoff()
	NotifyLandingApproved()
	NotifyTakeoffApproved()
}
