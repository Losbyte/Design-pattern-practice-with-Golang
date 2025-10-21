package newImp

import (
	"testing"
)


// TestCommodity 对应Java的test_commodity测试方法
func TestCommodity(t *testing.T) {
	// 创建工厂实例
	storeFactory := NewStoreFactory()

	// 1. 测试优惠券发放
	commodityService1, err := storeFactory.GetCommodityService(1)
	if err != nil {
		t.Fatalf("获取优惠券服务失败: %v", err)
	}
	err = commodityService1.SendCommodity("10001", "EGM1023938910232121323432", "791098764902132", nil)
	if err != nil {
		t.Errorf("优惠券发放失败: %v", err)
	}

	// 2. 测试实物商品发放
	commodityService2, err := storeFactory.GetCommodityService(2)
	if err != nil {
		t.Fatalf("获取实物商品服务失败: %v", err)
	}
	// 构建扩展信息map
	extMap := map[string]string{
		"consigneeUserName":  "谢飞机",
		"consigneeUserPhone": "15200292123",
		"consigneeUserAddress": "吉林省.长春市.双阳区.XX街道.檀溪苑小区.#18-2109",
	}
	err = commodityService2.SendCommodity("10001", "9820198721311", "1023000020112221113", extMap)
	if err != nil {
		t.Errorf("实物商品发放失败: %v", err)
	}

	// 3. 测试第三方兑换卡(爱奇艺)发放
	commodityService3, err := storeFactory.GetCommodityService(3)
	if err != nil {
		t.Fatalf("获取爱奇艺卡服务失败: %v", err)
	}
	err = commodityService3.SendCommodity("10001", "AQY1xjkUodl8LO975GdfrYUio", "", nil)
	if err != nil {
		t.Errorf("爱奇艺卡发放失败: %v", err)
	}
}
