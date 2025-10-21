package chain

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T){
	// 创建处理器
	auth := &AuthHandler{}
	logging := &LoggingHandler{}
	routing := &RoutingHandler{}
	cache := NewCacheHandler()

	// 构建处理链
	auth.SetNext(logging)
	logging.SetNext(cache)
	cache.SetNext(routing)

	// 创建请求
	validRequest := &Request{
		Method: "GET",
		Path:   "/api/users",
		Headers: map[string]string{
			"Authorization": "valid-token",
		},
	}

	invalidTokenRequest := &Request{
		Method: "GET",
		Path:   "/api/products",
		Headers: map[string]string{
			"Authorization": "invalid-token",
		},
	}

	noAuthRequest := &Request{
		Method:  "POST",
		Path:    "/api/users",
		Headers: map[string]string{},
	}

	// 处理请求
	fmt.Println("=== 处理有效请求 ===")
	auth.Handle(validRequest)

	fmt.Println("\n=== 处理无效token请求 ===")
	auth.Handle(invalidTokenRequest)

	fmt.Println("\n=== 处理无认证请求 ===")
	auth.Handle(noAuthRequest)

	fmt.Println("\n=== 第二次处理有效请求（测试缓存） ===")
	validRequest.Body = "User data" // 模拟第一次请求的结果
	auth.Handle(validRequest)

	fmt.Println("\n=== 责任链模式解决的问题 ===")
	fmt.Println("1. 解耦请求发送者和接收者")
	fmt.Println("2. 动态调整处理流程")
	fmt.Println("3. 多处理者协作")
	fmt.Println("4. 灵活分配责任")
}
