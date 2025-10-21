package newImp

// ICommodity 定义商品发放接口，对应Java的ICommodity接口
type ICommodity interface {
	// SendCommodity 发放商品
	// uId: 用户唯一标识
	// commodityId: 商品ID
	// bizId: 业务ID
	// extMap: 扩展信息
	// 返回error表示发放过程中出现的异常，对应Java的throws Exception
	SendCommodity(uId string, commodityId string, bizId string, extMap map[string]string) error
}
