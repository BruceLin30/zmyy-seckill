package main

import (
	"fmt"
	"zmyy_seckill/consts"
	"zmyy_seckill/zhimiaoyiyue"
)

func main() {
	//设置抢购请求速率，2s/次
	consts.RequestLimitRate.SetRate(1, 2)
	e := zhimiaoyiyue.ZMYYEngine{}
	e.Init()
	customerId := -1
	productId := -1
	for customerId == -1 || productId == -1 {
		if customerId == -1 {
			//获取指定地区接种地点的cutomerId
			customerId, _ = e.GetCustomerList()
		}
		if productId == -1 {
			//获取指定接种地点的productId
			productId, _ = e.GetCustomerProduct(customerId)
		}
	}
	//loc, _ := time.LoadLocation("Local")
	//timeLayout := "2006-01-02 15:04:05"
	//subsTime, _ := time.ParseInLocation(timeLayout, e.Conf.SubscribeTime, loc)
	//now := time.Now()
	//timer := time.NewTimer(subsTime.Sub(now))
	//<-timer.C
	e.Run(customerId, productId)
	fmt.Println("运行结束.....")
}
