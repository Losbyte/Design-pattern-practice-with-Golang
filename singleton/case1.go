package singleton

import "sync"

type singleton struct {

}

var (
	instance *singleton
	once sync.Once
)

func GetInstance() *singleton{
	once.Do(func(){
		instance = &singleton{}
	})
	return instance
}
