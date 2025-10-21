package goods

type DeliverReq struct {
	userName              string // 用户姓名
	userPhone             string // 用户手机
	sku                   string // 商品SKU
	orderId               string // 订单ID
	consigneeUserName     string // 收货人姓名
	consigneeUserPhone    string // 收货人手机
	consigneeUserAddress  string // 收货人地址
}

// GetUserName 获取用户姓名
func (d *DeliverReq) GetUserName() string {
	return d.userName
}

// SetUserName 设置用户姓名
func (d *DeliverReq) SetUserName(userName string) {
	d.userName = userName
}

// GetUserPhone 获取用户手机
func (d *DeliverReq) GetUserPhone() string {
	return d.userPhone
}

// SetUserPhone 设置用户手机
func (d *DeliverReq) SetUserPhone(userPhone string) {
	d.userPhone = userPhone
}

// GetSku 获取商品SKU
func (d *DeliverReq) GetSku() string {
	return d.sku
}

// SetSku 设置商品SKU
func (d *DeliverReq) SetSku(sku string) {
	d.sku = sku
}

// GetOrderId 获取订单ID
func (d *DeliverReq) GetOrderId() string {
	return d.orderId
}

// SetOrderId 设置订单ID
func (d *DeliverReq) SetOrderId(orderId string) {
	d.orderId = orderId
}

// GetConsigneeUserName 获取收货人姓名
func (d *DeliverReq) GetConsigneeUserName() string {
	return d.consigneeUserName
}

// SetConsigneeUserName 设置收货人姓名
func (d *DeliverReq) SetConsigneeUserName(consigneeUserName string) {
	d.consigneeUserName = consigneeUserName
}

// GetConsigneeUserPhone 获取收货人手机
func (d *DeliverReq) GetConsigneeUserPhone() string {
	return d.consigneeUserPhone
}

// SetConsigneeUserPhone 设置收货人手机
func (d *DeliverReq) SetConsigneeUserPhone(consigneeUserPhone string) {
	d.consigneeUserPhone = consigneeUserPhone
}

// GetConsigneeUserAddress 获取收货人地址
func (d *DeliverReq) GetConsigneeUserAddress() string {
	return d.consigneeUserAddress
}

// SetConsigneeUserAddress 设置收货人地址
func (d *DeliverReq) SetConsigneeUserAddress(consigneeUserAddress string) {
	d.consigneeUserAddress = consigneeUserAddress
}
