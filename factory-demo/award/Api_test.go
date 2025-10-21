package award

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestAwardToUser(t *testing.T) {
	prizeController := &PrizeController{}
	fmt.Println("模拟发放优惠券测试")
	// 模拟发放优惠券测试
	req01 := &AwardReq{}
	req01.SetUId("10001")
	req01.SetAwardType(1)
	req01.SetAwardNumber("EGM1023938910232121323432")
	req01.SetBizId("791098764902132")
	awardRes01 := prizeController.AwardToUser(req01)
	req01Json, _ := json.Marshal(req01)
	res01Json, _ := json.Marshal(awardRes01)
	log.Printf("请求参数：%s", req01Json)
	log.Printf("测试结果：%s", res01Json)
	fmt.Println("模拟发放实物商品")
	// 模拟发放实物商品
	req02 := &AwardReq{}
	req02.SetUId("10001")
	req02.SetAwardType(2)
	req02.SetAwardNumber("9820198721311")
	req02.SetBizId("1023000020112221113")
	extMap := map[string]string{
		"consigneeUserName":  "谢飞机",
		"consigneeUserPhone": "15200292123",
		"consigneeUserAddress": "吉林省.长春市.双阳区.XX街道.檀溪苑小区.#18-2109",
	}
	req02.SetExtMap(extMap)
	awardRes02 := prizeController.AwardToUser(req02)
	req02Json, _ := json.Marshal(req02)
	res02Json, _ := json.Marshal(awardRes02)
	log.Printf("请求参数：%s", req02Json)
	log.Printf("测试结果：%s", res02Json)
	fmt.Println("第三方兑换卡(爱奇艺)")
	req03 := &AwardReq{}
	req03.SetUId("10001")
	req03.SetAwardType(3)
	req03.SetAwardNumber("AQY1xjkUodl8LO975GdfrYUio")
	awardRes03 := prizeController.AwardToUser(req03)
	req03Json, _ := json.Marshal(req03)
	res03Json, _ := json.Marshal(awardRes03)
	log.Printf("请求参数：%s", req03Json)
	log.Printf("测试结果：%s", res03Json)
}
