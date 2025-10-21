package award

// AwardReq 奖品请求信息结构体，对应Java的AwardReq类
type AwardReq struct {
	uId         string            // 用户唯一ID
	awardType   int               // 奖品类型(可以用枚举定义)；1优惠券、2实物商品、3第三方兑换卡(爱奇艺)
	awardNumber string            // 奖品编号；sku、couponNumber、cardId
	bizId       string            // 业务ID，防重复
	extMap      map[string]string // 扩展信息
}

// GetUId 获取用户唯一ID，对应Java的getuId方法
func (a *AwardReq) GetUId() string {
	return a.uId
}

// SetUId 设置用户唯一ID，对应Java的setuId方法
func (a *AwardReq) SetUId(uId string) {
	a.uId = uId
}

// GetAwardType 获取奖品类型，对应Java的getAwardType方法
func (a *AwardReq) GetAwardType() int {
	return a.awardType
}

// SetAwardType 设置奖品类型，对应Java的setAwardType方法
func (a *AwardReq) SetAwardType(awardType int) {
	a.awardType = awardType
}

// GetAwardNumber 获取奖品编号，对应Java的getAwardNumber方法
func (a *AwardReq) GetAwardNumber() string {
	return a.awardNumber
}

// SetAwardNumber 设置奖品编号，对应Java的setAwardNumber方法
func (a *AwardReq) SetAwardNumber(awardNumber string) {
	a.awardNumber = awardNumber
}

// GetBizId 获取业务ID，对应Java的getBizId方法
func (a *AwardReq) GetBizId() string {
	return a.bizId
}

// SetBizId 设置业务ID，对应Java的setBizId方法
func (a *AwardReq) SetBizId(bizId string) {
	a.bizId = bizId
}

// GetExtMap 获取扩展信息，对应Java的getExtMap方法
func (a *AwardReq) GetExtMap() map[string]string {
	return a.extMap
}

// SetExtMap 设置扩展信息，对应Java的setExtMap方法
func (a *AwardReq) SetExtMap(extMap map[string]string) {
	a.extMap = extMap
}
