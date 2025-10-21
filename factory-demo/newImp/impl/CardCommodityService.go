package impl

import (
	"Demo1/card"
	"encoding/json"
	"log"
)


// CardCommodityService 实现ICommodity接口，处理爱奇艺兑换卡发放
type CardCommodityService struct {
	iQiYiCardService card.IQiYiCardService // 爱奇艺卡服务
}

// NewCardCommodityService 构造函数
func NewCardCommodityService() *CardCommodityService {
	return &CardCommodityService{
		iQiYiCardService: card.NewIQiYiCardService(), // 模拟注入
	}
}

// SendCommodity 实现ICommodity接口的发放商品方法
func (c *CardCommodityService) SendCommodity(uId string, commodityId string, bizId string, extMap map[string]string) error {
	mobile := c.queryUserMobile(uId)
	c.iQiYiCardService.GrantToken(mobile, bizId)

	// 序列化extMap用于日志输出
	extMapJson, err := json.Marshal(extMap)
	if err != nil {
		extMapJson = []byte("{}") // 序列化失败时使用空对象
	}

	log.Printf("请求参数[爱奇艺兑换卡] => uId：%s commodityId：%s bizId：%s extMap：%s",
		uId, commodityId, bizId, extMapJson)
	log.Println("测试结果[爱奇艺兑换卡]：success")

	return nil // 无错误返回nil，对应Java的无异常抛出
}

// queryUserMobile 查询用户手机号（私有方法）
func (c *CardCommodityService) queryUserMobile(uId string) string {
	return "15200101232" // 模拟返回固定手机号
}
