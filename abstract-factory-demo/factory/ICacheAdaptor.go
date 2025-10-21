package factory

import "time"

// ICacheAdapter 缓存适配器接口，定义缓存操作规范
type ICacheAdapter interface {
	// Get 根据key获取缓存值
	Get(key string) string
	// Set 设置缓存键值对（无过期时间）
	Set(key, value string)
	// SetWithTimeout 设置缓存键值对（带过期时间）
	SetWithTimeout(key, value string, timeout time.Duration)
	// Del 删除指定缓存键
	Del(key string)
}
