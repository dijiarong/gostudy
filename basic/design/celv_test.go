package design

import (
	"fmt"
	"testing"
)

type Send interface {
	SendMsg()
}

type SendMsger interface {
	sendFeiShuMsg()
	sendDingDingMsg()
	sendEmailMsg()
	sendWeChatMsg()
	sendMobileMsg()
}

type BaseMaiMai struct {
	msg string
	SendMsger
}

func (fs MaiMai) SendMsg() {
	fs.sendFeiShuMsg()
	fs.sendDingDingMsg()
	fs.sendEmailMsg()
	fs.sendWeChatMsg()
	fs.sendMobileMsg()
}

type MaiMai struct {
	BaseMaiMai
}

func NewMaiMai(msg string) Send {
	maiMai := new(MaiMai)
	maiMai.msg = msg
	maiMai.SendMsger = maiMai
	return maiMai
}

func (mm *MaiMai) sendFeiShuMsg() {
	fmt.Printf("飞书发送了: %s\n", mm.msg)
}

func (mm *MaiMai) sendDingDingMsg() {
	fmt.Printf("丁丁发送了: %s\n", mm.msg)
}
func (mm *MaiMai) sendEmailMsg() {
	fmt.Printf("邮件发送了: %s\n", mm.msg)
}
func (mm *MaiMai) sendWeChatMsg() {
	fmt.Printf("微信发送了: %s\n", mm.msg)
}
func (mm *MaiMai) sendMobileMsg() {
	fmt.Printf("手机短信发送了: %s\n", mm.msg)
}

func TestCelueMethod(t *testing.T) {
	NewMaiMai("你好").SendMsg()
}
