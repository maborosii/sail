package sender

import (
	cm "sail/common"
	"sail/global"
	"sync"

	"go.uber.org/zap"
)

type Pusher interface {
	// push action
	Push(m cm.OutMessage) error
	// pusher type
	Type() string
}

// just for test
func PushMessage(p Pusher, m cm.OutMessage) error {
	p.Push(m)
	return nil
}

type PushList struct {
	Pushers []Pusher
}

// 初始化
func NewPusherList() *PushList {
	return &PushList{}
}

// 注册pushers
func (p *PushList) RegisterPusher(senders ...Pusher) {
	p.Pushers = append(p.Pushers, senders...)
}

// 批量执行推送任务
func (p *PushList) Exec(om cm.OutMessage) {
	w := sync.WaitGroup{}
	for _, pp := range p.Pushers {
		w.Add(1)
		go func(om cm.OutMessage, pp Pusher) {
			defer w.Done()
			err := pp.Push(om)
			global.Logger.Error("it occurs error when push messgae", zap.Error(err))
		}(om, pp)
	}
	w.Wait()
}
