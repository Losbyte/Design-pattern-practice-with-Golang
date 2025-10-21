package newImp

import (
	"Demo1/newImp/impl"
	"errors"
)


// StoreFactory 商品服务工厂类，用于创建不同类型的商品服务实例
type StoreFactory struct{}

// NewStoreFactory 创建StoreFactory实例
func NewStoreFactory() *StoreFactory {
	return &StoreFactory{}
}

// GetCommodityService 根据商品类型获取对应的商品服务实例
// commodityType: 商品类型 1-优惠券 2-实物商品 3-卡券
// 返回实现了store.ICommodity接口的实例，以及可能的错误
func (f *StoreFactory) GetCommodityService(commodityType int) (ICommodity, error) {
	if commodityType <= 0 {
		return nil, errors.New("商品类型不能为负数或零")
	}

	switch commodityType {
	case 1:
		return impl.NewCouponCommodityService(), nil
	case 2:
		return impl.NewGoodsCommodityService(), nil
	case 3:
		return impl.NewCardCommodityService(), nil
	default:
		return nil, errors.New("不存在的商品服务类型")
	}
}
