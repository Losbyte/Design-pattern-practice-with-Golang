package observer

import "testing"

func TestObserver(t *testing.T)  {
	//  创建主题对象
	subject  :=  &Subject{}
	//  创建具体观察者对象
	observerA  :=  &ConcreteObserverA{name:  "Observer  A"}

	//  将观察者添加到主题的观察者列表中
	subject.Attach(observerA)

	//  当主题状态改变时，通知所有观察者
	subject.Notify("State  changed  to  State  1")
}
