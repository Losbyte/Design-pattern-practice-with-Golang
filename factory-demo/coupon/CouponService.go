package coupon

import "fmt"

// CouponService 模拟优惠券服务
type CouponService struct{}

func NewCouponService()CouponService{
	return CouponService{}
}

// SendCoupon 发放优惠券
// uId: 用户ID
// couponNumber: 优惠券编号
// uuid: 唯一标识
func (c *CouponService) SendCoupon(uId, couponNumber, uuid string) *CouponResult {
	fmt.Printf("模拟发放优惠券一张：%s,%s,%s\n", uId, couponNumber, uuid)
	return NewCouponResult("0000", "发放成功")
}
