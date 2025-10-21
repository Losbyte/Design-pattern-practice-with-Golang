package award

import (
	"encoding/json"
	"fmt"
	"log"
	"Demo1/card"
	"Demo1/coupon"
	"Demo1/goods"
)

// PrizeController 模拟发奖控制器
type PrizeController struct{}

// AwardToUser 处理用户奖品发放
func (p *PrizeController) AwardToUser(req *AwardReq) *AwardRes {
	// 序列化请求参数
	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Printf("JSON序列化失败: %v", err)
		return NewAwardRes("0001", "请求参数序列化失败")
	}

	var awardRes *AwardRes
	// 使用defer+recover捕获panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("奖品发放panic异常 uId:%s, req:%s, err:%v", req.GetUId(), reqJson, r)
			awardRes = NewAwardRes("0001", "系统异常")
		}
	}()

	// 业务逻辑处理（Go中使用显式错误判断替代try-catch）
	log.Printf("奖品发放开始 %s。req:%s", req.GetUId(), reqJson)

	switch req.GetAwardType() {
	case 1:
		// 优惠券发放
		couponService := &coupon.CouponService{}
		couponResult := couponService.SendCoupon(req.GetUId(), req.GetAwardNumber(), req.GetBizId())
		if couponResult.GetCode() == "0000" {
			awardRes = NewAwardRes("0000", "发放成功")
		} else {
			awardRes = NewAwardRes("0001", couponResult.GetInfo())
		}
	case 2:
		// 实物商品发放
		goodsService := &goods.GoodsService{}
		deliverReq := &goods.DeliverReq{}
		deliverReq.SetUserName(p.queryUserName(req.GetUId()))
		deliverReq.SetUserPhone(p.queryUserPhoneNumber(req.GetUId()))
		deliverReq.SetSku(req.GetAwardNumber())
		deliverReq.SetOrderId(req.GetBizId())

		extMap := req.GetExtMap()
		if extMap != nil {
			deliverReq.SetConsigneeUserName(extMap["consigneeUserName"])
			deliverReq.SetConsigneeUserPhone(extMap["consigneeUserPhone"])
			deliverReq.SetConsigneeUserAddress(extMap["consigneeUserAddress"])
		}

		isSuccess := goodsService.DeliverGoods(deliverReq)
		if isSuccess {
			awardRes = NewAwardRes("0000", "发放成功")
		} else {
			awardRes = NewAwardRes("0001", "发放失败")
		}
	case 3:
		// 第三方兑换卡发放
		bindMobileNumber := p.queryUserPhoneNumber(req.GetUId())
		iQiYiCardService := &card.IQiYiCardService{}
		iQiYiCardService.GrantToken(bindMobileNumber, req.GetAwardNumber())
	default:
		awardRes = NewAwardRes("0001", fmt.Sprintf("不支持的奖品类型: %d", req.GetAwardType()))
	}

	log.Printf("奖品发放完成 %s。", req.GetUId())
	return awardRes
}

// queryUserName 查询用户姓名
func (p *PrizeController) queryUserName(uId string) string {
	return "花花"
}

// queryUserPhoneNumber 查询用户手机号
func (p *PrizeController) queryUserPhoneNumber(uId string) string {
	return "15200101232"
}

