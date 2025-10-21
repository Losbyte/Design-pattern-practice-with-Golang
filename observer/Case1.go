package observer

import (
	"fmt"
)

type  Observer  interface  {
  Update(string)  //  当主题状态改变时，此方法会被调用
}

//  定义Subject结构体，它包含一个观察者列表和方法来添加或通知观察者
type  Subject  struct  {
  observers  []Observer  //  存储观察者的列表
}

//  Attach方法用于将一个观察者添加到观察者列表中
func  (s  *Subject)  Attach(observer  Observer)  {
	  s.observers  =  append(s.observers,  observer)
}

//  Notify方法用于通知所有观察者主题状态的改变
func  (s  *Subject)  Notify(message  string)  {
	  for  _,  observer  :=  range  s.observers  {
		          observer.Update(message)  //  调用每个观察者的Update方法
		  }
}

//  定义一个具体观察者ConcreteObserver，它实现了Observer接口
type  ConcreteObserverA  struct  {
  name  string
}

//  实现Observer接口的Update方法
func  (c  *ConcreteObserverA)  Update(message  string)  {
	  fmt.Printf("%s  received  message:  %s\n",  c.name,  message)
}


