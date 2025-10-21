package factory

import "time"

// JDKProxy 模拟Java的JDKProxy类，用于创建代理对象
type GoProxy struct{}

// GetProxy 生成实现指定接口的代理对象
// 参数：
//   - interfaceType: 目标接口类型（需传入接口的反射Type）
//   - cacheAdapter: 实际的缓存适配器实现
// 返回：
//   - 实现了目标接口的代理对象
func (j *GoProxy) GetProxy(cacheAdapter ICacheAdapter) (interface{}, error) {
	// 创建调用处理器
	handler := NewGoInvocationHandler(cacheAdapter)

	// 创建代理对象（Go中通过自定义结构体实现代理）
	proxy := &proxyObject{
		handler: handler,
	}
	return proxy, nil
}

// proxyObject 代理对象结构体，实现目标接口
type proxyObject struct {
	handler *GoInvocationHandler // 调用处理器
}

// 实现目标接口的方法（以Get为例，其他方法类似实现）
func (p *proxyObject) Get(key string) string {
	result, _ := p.handler.Invoke("Get", key)
	return result.(string)
}

// Set 实现目标接口的Set方法
func (p *proxyObject) Set(key, value string) {
	p.handler.Invoke("Set", key, value)
}

// SetWithTimeout 实现目标接口的SetWithTimeout方法
func (p *proxyObject) SetWithTimeout(key, value string, timeout time.Duration) {
	p.handler.Invoke("SetWithTimeout", key, value, timeout)
}

// Del 实现目标接口的Del方法
func (p *proxyObject) Del(key string) {
	p.handler.Invoke("Del", key)
}

