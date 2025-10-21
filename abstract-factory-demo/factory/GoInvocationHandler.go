package factory

import (
	"fmt"
	"reflect"
)

type GoInvocationHandler struct {
	cacheAdapter ICacheAdapter // 持有缓存适配器实例
}

func NewGoInvocationHandler(adapter ICacheAdapter) *GoInvocationHandler {
	return &GoInvocationHandler{
		cacheAdapter: adapter,
	}
}

func (h *GoInvocationHandler) Invoke(methodName string, args ...interface{}) (interface{}, error) {
	// 获取缓存适配器的反射类型
	adapterType := reflect.TypeOf(h.cacheAdapter)
	// 查找对应的方法
	method, ok := adapterType.MethodByName(methodName)
	if !ok {
		return nil, fmt.Errorf("方法 %s 不存在", methodName)
	}
	// 调用方法并返回结果
	result := method.Func.Call(append([]reflect.Value{reflect.ValueOf(h.cacheAdapter)}, argsToValues(args)...))
	// 处理返回值（如果有）
	if len(result) > 0 {
		return result[0].Interface(), nil
	}
	return nil, nil
}

// argsToValues 将参数转换为reflect.Value切片
func argsToValues(args []interface{}) []reflect.Value {
	values := make([]reflect.Value, len(args))
	for i, arg := range args {
		values[i] = reflect.ValueOf(arg)
	}
	return values
}
