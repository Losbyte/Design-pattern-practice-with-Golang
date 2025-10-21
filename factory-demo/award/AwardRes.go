package award

// AwardRes 对应Java的AwardRes类，用于表示奖品发放结果
type AwardRes struct {
	code string // 编码
	info string // 描述
}

// NewAwardRes 构造函数，对应Java的构造方法
func NewAwardRes(code, info string) *AwardRes {
	return &AwardRes{
		code: code,
		info: info,
	}
}

// GetCode 获取编码，对应Java的getCode方法
func (a *AwardRes) GetCode() string {
	return a.code
}

// SetCode 设置编码，对应Java的setCode方法
func (a *AwardRes) SetCode(code string) {
	a.code = code
}

// GetInfo 获取描述，对应Java的getInfo方法
func (a *AwardRes) GetInfo() string {
	return a.info
}

// SetInfo 设置描述，对应Java的setInfo方法
func (a *AwardRes) SetInfo(info string) {
	a.info = info
}
