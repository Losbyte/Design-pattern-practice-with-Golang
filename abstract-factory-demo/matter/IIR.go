package matter

import (
	"log"
	"sync"
	"time"
)


// IIR 模拟IIR缓存实现
type IIR struct {
	dataMap sync.Map // 并发安全的map，对应Java的ConcurrentHashMap
}

// NewIIR 创建IIR实例
func NewIIR() *IIR {
	return &IIR{
		dataMap: sync.Map{},
	}
}

// Get 根据key获取数据
func (i *IIR) Get(key string) string {
	log.Printf("IIR获取数据 key：%s", key)
	val, ok := i.dataMap.Load(key)
	if ok {
		return val.(string)
	}
	return ""
}

// Set 写入数据（无过期时间）
func (i *IIR) Set(key, value string) {
	log.Printf("IIR写入数据 key：%s val：%s", key, value)
	i.dataMap.Store(key, value)
}

// SetExpire 写入数据（带过期时间）
func (i *IIR) SetExpire(key, value string, timeout int64, timeUnit time.Duration) {
	log.Printf("IIR写入数据 key：%s val：%s timeout：%d timeUnit：%s", key, value, timeout, timeUnit.String())
	i.dataMap.Store(key, value)
	// 实际业务中可根据需要添加过期清理逻辑，例如：
	// time.AfterFunc(time.Duration(timeout)*timeUnit, func() {
	//     i.dataMap.Delete(key)
	// })
}

// Del 删除数据
func (i *IIR) Del(key string) {
	log.Printf("IIR删除数据 key：%s", key)
	i.dataMap.Delete(key)
}
