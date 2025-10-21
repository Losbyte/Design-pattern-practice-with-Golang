package impl

import (
	"Demo1/goods"
	"encoding/json"
	"fmt"
	"log"
)


// GoodsCommodityService 实现ICommodity接口，处理实物商品发放
type GoodsCommodityService struct {
	goodsService goods.GoodsService // 商品服务
}

// NewGoodsCommodityService 构造函数
func NewGoodsCommodityService() *GoodsCommodityService {
	return &GoodsCommodityService{
		goodsService: goods.NewGoodsService(), // 初始化商品服务
	}
}

// SendCommodity 实现ICommodity接口的发放商品方法
func (g *GoodsCommodityService) SendCommodity(uId string, commodityId string, bizId string, extMap map[string]string) error {
	// 构建发货请求
	deliverReq := &goods.DeliverReq{}
	deliverReq.SetUserName(g.queryUserName(uId))
	deliverReq.SetUserPhone(g.queryUserPhoneNumber(uId))
	deliverReq.SetSku(commodityId)
	deliverReq.SetOrderId(bizId)
	deliverReq.SetConsigneeUserName(extMap["consigneeUserName"])
	deliverReq.SetConsigneeUserPhone(extMap["consigneeUserPhone"])
	deliverReq.SetConsigneeUserAddress(extMap["consigneeUserAddress"])
	// 调用商品服务发货
	isSuccess := g.goodsService.DeliverGoods(deliverReq)

	// 序列化扩展信息用于日志
	extMapJson, err := json.Marshal(extMap)
	if err != nil {
		extMapJson = []byte("{}")
	}

	// 打印请求参数和结果日志
	log.Printf("请求参数[实物商品] => uId：%s commodityId：%s bizId：%s extMap：%s",
		uId, commodityId, bizId, extMapJson)
	log.Printf("测试结果[实物商品]：%v", isSuccess)

	// 发货失败时返回错误
	if !isSuccess {
		return fmt.Errorf("发货失败") // 自定义错误类型
	}

	return nil
}

// 查询用户名（私有方法）
func (g *GoodsCommodityService) queryUserName(uId string) string {
	return "花花"
}

// 查询用户手机号（私有方法）
func (g *GoodsCommodityService) queryUserPhoneNumber(uId string) string {
	return "15200101232"
}
