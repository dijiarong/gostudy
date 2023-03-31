package design

import (
	"fmt"
	"sync"
	"testing"
)

// 演员接口
type IActor interface {
	DressUp() string
}

type dressBehavior interface {
	makeUp() string
	clothe() string
	wear() string
}

// 演员的基类
type BaseActor struct {
	roleName string // 角色名称
	dressBehavior
}

func (b *BaseActor) DressUp() string {
	return "扮演" + b.roleName + "的" + b.makeUp() + b.clothe() + b.wear()
}

// 具体演员
type womanActor struct {
	BaseActor
}

func NewWomanActor(roleName string) *womanActor {
	actor := new(womanActor)
	actor.roleName = roleName
	actor.dressBehavior = actor
	return actor
}

// 化妆
func (w *womanActor) makeUp() string {
	return "女演员涂着口红，画着眉毛；"
}

// 穿衣
func (w *womanActor) clothe() string {
	return "穿着连衣裙；"
}

var WomanActorGobal womanActor

func TestTemplateMethod(t *testing.T) {
	one := sync.Once{}
	one.Do(func() {
		WomanActorGobal = womanActor{}
	})
	showActors(NewWomanActor("妈妈"))
}

// showActors 显示演员的装扮信息
func showActors(actors ...IActor) {
	for _, actor := range actors {
		fmt.Println(actor.DressUp())
	}
}
