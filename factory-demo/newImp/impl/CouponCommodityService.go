package impl

import (
	"Demo1/coupon"
	"encoding/json"
	"fmt"
	"log"
)


// CouponCommodityService 实现ICommodity接口，处理优惠券发放
type CouponCommodityService struct {
	couponService coupon.CouponService // 优惠券服务
}

// NewCouponCommodityService 构造函数
func NewCouponCommodityService() *CouponCommodityService {
	return &CouponCommodityService{
		couponService: coupon.NewCouponService(), // 初始化优惠券服务
	}
}

// SendCommodity 实现ICommodity接口的发放商品方法
func (c *CouponCommodityService) SendCommodity(uId string, commodityId string, bizId string, extMap map[string]string) error {
	// 调用优惠券服务发放优惠券
	couponResult := c.couponService.SendCoupon(uId, commodityId, bizId)

	// 序列化扩展信息用于日志输出
	extMapJson, err := json.Marshal(extMap)
	if err != nil {
		extMapJson = []byte("{}") // 序列化失败时使用空对象
	}

	// 序列化返回结果用于日志输出
	resultJson, err := json.Marshal(couponResult)
	if err != nil {
		resultJson = []byte("{}")
	}

	// 打印请求参数和测试结果日志
	log.Printf("请求参数[优惠券] => uId：%s commodityId：%s bizId：%s extMap：%s",
		uId, commodityId, bizId, extMapJson)
	log.Printf("测试结果[优惠券]：%s", resultJson)

	// 校验返回状态，非成功状态返回错误
	if couponResult.GetCode() != "0000" {
		return fmt.Errorf(couponResult.GetInfo())
	}

	return nil // 成功返回nil
}
