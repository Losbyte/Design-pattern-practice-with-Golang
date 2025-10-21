package impl

import (
	"Demo2/matter"
	"time"
)

// IIRCacheAdapter 实现ICacheAdapter接口，适配IIR缓存
type IIRCacheAdapter struct {
	iir *matter.IIR // 持有IIR实例，对应Java的private IIR iir = new IIR()
}

// NewIIRCacheAdapter 创建IIRCacheAdapter实例
func NewIIRCacheAdapter() *IIRCacheAdapter {
	return &IIRCacheAdapter{
		iir: matter.NewIIR(), // 初始化IIR实例
	}
}

// Get 实现ICacheAdapter接口的Get方法
func (i *IIRCacheAdapter) Get(key string) string {
	return i.iir.Get(key) // 调用IIR的Get方法，对应Java的iir.get(key)
}

// Set 实现ICacheAdapter接口的Set方法（无过期时间）
func (i *IIRCacheAdapter) Set(key, value string) {
	i.iir.Set(key, value) // 调用IIR的Set方法，对应Java的iir.set(key, value)
}

// SetWithTimeout 实现ICacheAdapter接口的Set方法（带过期时间）
// Go中不支持方法重载，用不同方法名区分
func (i *IIRCacheAdapter) SetWithTimeout(key, value string, timeout time.Duration) {
	// 将time.Duration转换为基础时间单位和数值，适配IIR的SetExpire方法
	// 对应Java的iir.setExpire(key, value, timeout, timeUnit)
	// 这里以秒为单位拆分duration（可根据实际需求调整）
	seconds := int64(timeout.Seconds())
	i.iir.SetExpire(key, value, seconds, time.Second)
}

// Del 实现ICacheAdapter接口的Del方法
func (i *IIRCacheAdapter) Del(key string) {
	i.iir.Del(key) // 调用IIR的Del方法，对应Java的iir.del(key)
}
