package coupon

type CouponResult struct {
	code string // 编码
	info string // 描述
}

// NewCouponResult 创建CouponResult实例
func NewCouponResult(code, info string) *CouponResult {
	return &CouponResult{
		code: code,
		info: info,
	}
}

// GetCode 获取编码
func (c *CouponResult) GetCode() string {
	return c.code
}

// SetCode 设置编码
func (c *CouponResult) SetCode(code string) {
	c.code = code
}

// GetInfo 获取描述
func (c *CouponResult) GetInfo() string {
	return c.info
}

// SetInfo 设置描述
func (c *CouponResult) SetInfo(info string) {
	c.info = info
}
