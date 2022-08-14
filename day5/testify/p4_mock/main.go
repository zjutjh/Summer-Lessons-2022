package main

import (
	"fmt"
)

// MessageService 通知客户被收取的费用
type MessageService interface {
	SendChargeNotification(int) bool
}

// SMSService 是 MessageService 的实现
type SMSService struct{}

// MyService 使用 MessageService 来通知客户
type MyService struct {
	messageService MessageService
}

// SendChargeNotification 通过 SMS 来告知客户他们被收取费用
// 这就是我们将要模拟的方法
func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending Production Charge Notification")
	return true
}

// ChargeCustomer 向客户收取费用
// 在真实系统中，我们会模拟这个
// 但是在这里，我想在每次运行测试时都赚点钱
func (a MyService) ChargeCustomer(value int) error {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil
}

func main() {
	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}
