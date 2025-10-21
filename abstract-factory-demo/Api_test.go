package abstract_factory_demo

import (
	"Demo2/factory"
	"Demo2/factory/impl"
	"fmt"
	"testing"
)

func TestCacheService(t *testing.T){
	// 测试EGM缓存适配器
	jdkProxy := &factory.GoProxy{}
	egmAdapter := impl.NewEGMCacheAdapter()
	proxyEGM, err := jdkProxy.GetProxy(egmAdapter)
	if err != nil {
		t.Fatalf("创建EGM代理失败: %v", err)
	}
	cache := proxyEGM.(factory.ICacheAdapter)
	cache.Set("user_name_01", "小傅哥")
	val01 := cache.Get("user_name_01")
	fmt.Println("EGM测试结果：" + val01)

	// 测试IIR缓存适配器
	jdkProxy = &factory.GoProxy{}
	iirAdapter := impl.NewIIRCacheAdapter()
	proxyIIR, err := jdkProxy.GetProxy(iirAdapter)
	if err != nil {
		t.Fatalf("创建IIR代理失败: %v", err)
	}
	cache = proxyIIR.(factory.ICacheAdapter)
	cache.Set("user_name_01", "小傅哥")
	val02 := cache.Get("user_name_01")
	fmt.Println("IIR测试结果：" + val02)
}
