package observer

import "testing"

func TestObserver(t *testing.T) {

	// 创建信用卡
	card := NewCreditCard("张三")
	// 短信通知定于信用卡消费和逾期信息
	card.Subscribe(new(shortMessage), ConsumeType, ExpireType)
	// 电子邮件通知订阅信用卡账单及逾期消息
	card.Subscribe(new(email), BillType, ExpireType)
	// 电话通知订阅信用卡逾期消息，同时逾期消息通过三种方式通知
	card.Subscribe(new(telephone), ExpireType)

	card.Consume(500.00)
	card.Consume(800.00)
	card.SendBill()
	card.Expire()
}
