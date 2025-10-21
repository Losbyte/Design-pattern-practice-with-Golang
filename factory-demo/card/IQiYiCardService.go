package card

import "fmt"

type IQiYiCardService struct {}


func NewIQiYiCardService() IQiYiCardService{
	return IQiYiCardService{}
}

func (q *IQiYiCardService) GrantToken(bindMobileNumber,cardId string){
	fmt.Printf("模拟发放爱奇艺会员卡一张：" + bindMobileNumber + "，" + cardId)
}

