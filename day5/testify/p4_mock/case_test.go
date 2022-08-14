package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type smsServiceMock struct {
	mock.Mock
}

// 我们模拟的 smsService 方法
func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed in: %d\n", value)
	// 这将记录方法被调用以及被调用时传进来的参数值
	args := m.Called(value)
	// 它将返回任何我需要返回的
	// 这种情况下模拟一个 SMS Service Notification 被发送出去
	return args.Bool(0)
}

// 我们将实现 MessageService 接口
// 这就意味着我们不得不改写在接口中定义的所有方法
func (m *smsServiceMock) DummyFunc() {
	fmt.Println("Dummy")
}

// TestChargeCustomer 是个奇迹发生的地方
// 在这里我们将创建 SMSService mock
func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)

	// 然后我们将定义当 100 传递给 SendChargeNotification 时，需要返回什么
	// 在这里，我们希望它在成功发送通知后返回 true
	smsService.On("SendChargeNotification", 100).Return(true)

	// 接下来，我们要定义要测试的服务
	myService := MyService{smsService}
	// 然后调用方法
	myService.ChargeCustomer(100)

	// 最后，我们验证 myService.ChargeCustomer 调用了我们模拟的 SendChargeNotification 方法
	smsService.AssertExpectations(t)
}
