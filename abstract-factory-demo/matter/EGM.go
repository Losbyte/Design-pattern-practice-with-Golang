package matter

import (
	"log"
	"sync"
	"time"
)

// EGM 模拟缓存实现
type EGM struct {
	dataMap sync.Map // 并发安全的map，对应Java的ConcurrentHashMap
}

// NewEGM 创建EGM实例
func NewEGM() *EGM {
	return &EGM{
		dataMap: sync.Map{},
	}
}

// Gain 根据key获取数据
func (e *EGM) Gain(key string) string {
	log.Printf("EGM获取数据 key：%s", key)
	val, ok := e.dataMap.Load(key)
	if ok {
		return val.(string)
	}
	return ""
}

// Set 写入数据（无过期时间）
func (e *EGM) Set(key, value string) {
	log.Printf("EGM写入数据 key：%s val：%s", key, value)
	e.dataMap.Store(key, value)
}

// SetEx 写入数据（带过期时间）
func (e *EGM) SetEx(key, value string, timeout int64, timeUnit time.Duration) {
	// 注意：Go中time.Duration已包含单位，这里参数对应Java的TimeUnit
	log.Printf("EGM写入数据 key：%s val：%s timeout：%d timeUnit：%s", key, value, timeout, timeUnit.String())
	e.dataMap.Store(key, value)
	// 实际业务中可根据需要添加过期清理逻辑
}

// Delete 删除数据
func (e *EGM) Delete(key string) {
	log.Printf("EGM删除数据 key：%s", key)
	e.dataMap.Delete(key)
}
