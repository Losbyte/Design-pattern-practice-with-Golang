package impl

import (
	"Demo2/matter"
	"time"
)


// EGMCacheAdapter 实现ICacheAdapter接口，适配EGM缓存
type EGMCacheAdapter struct {
	egm *matter.EGM // 持有EGM实例，对应Java的private EGM egm = new EGM()
}

// NewEGMCacheAdapter 创建EGMCacheAdapter实例
func NewEGMCacheAdapter() *EGMCacheAdapter {
	return &EGMCacheAdapter{
		egm: matter.NewEGM(), // 初始化EGM实例
	}
}

// Get 实现ICacheAdapter接口的Get方法，对应Java的get
func (e *EGMCacheAdapter) Get(key string) string {
	return e.egm.Gain(key) // 调用EGM的Gain方法，对应Java的egm.gain(key)
}

// Set 实现ICacheAdapter接口的Set方法（无过期时间），对应Java的set
func (e *EGMCacheAdapter) Set(key, value string) {
	e.egm.Set(key, value) // 调用EGM的Set方法，对应Java的egm.set(key, value)
}

// SetWithTimeout 实现ICacheAdapter接口的Set方法（带过期时间），对应Java的重载set
func (e *EGMCacheAdapter) SetWithTimeout(key, value string, timeout time.Duration) {
	// 将time.Duration转换为秒数和时间单位，适配EGM的SetEx方法
	// 对应Java的egm.setEx(key, value, timeout, timeUnit)
	// 注：这里假设EGM的SetEx参数为(key, value string, timeout int64, unit time.Duration)
	e.egm.SetEx(key, value, int64(timeout.Seconds()), time.Second)
}

// Del 实现ICacheAdapter接口的Del方法，对应Java的del
func (e *EGMCacheAdapter) Del(key string) {
	e.egm.Delete(key) // 调用EGM的Delete方法，对应Java的egm.delete(key)
}
