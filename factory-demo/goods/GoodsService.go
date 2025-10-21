package goods

import "encoding/json"
import "fmt"

// GoodsService 模拟实物商品服务
type GoodsService struct{}

func NewGoodsService()GoodsService{
	return GoodsService{}
}
// DeliverGoods 发货实物商品
// req: 配送请求信息
func (g *GoodsService) DeliverGoods(req *DeliverReq) bool {
	// 将请求信息转为JSON字符串，模拟原Java中的JSON.toJSONString
	reqJson, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("JSON序列化失败: %v\n", err)
		return false
	}
	fmt.Printf("模拟发货实物商品一个：%s\n", reqJson)
	return true
}
