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
	if err := p.Push(m); err != nil {
		return err
	}
	return nil
}

type PushList struct {
	Pushers []Pusher
}

// 初始化
func NewPusherList() *PushList {
	pl := make([]Pusher, 0)
	return &PushList{
		Pushers: pl,
	}
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
			defer func() {
				if err := recover(); err != nil {
					global.Logger.Error("push action in pusherlist occured error", zap.Any("error", err))
				}
			}()
			defer w.Done()
			err := pp.Push(om)
			if err != nil {
				global.Logger.Error("it occurs error when push messgae", zap.Error(err))
			}
		}(om, pp)
	}
	w.Wait()
}
