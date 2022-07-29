package controller

import (
	cm "sail/common"
	"sail/global"
	"sail/internal/sender"
	"sync"

	"go.uber.org/zap"
)

// 消息推送器队列
type PushList struct {
	Pushers []sender.Pusher
}

// 初始化
func (p *PushList) Init(senders ...sender.Pusher) {
	p.Pushers = make([]sender.Pusher, 0)
	p.Pushers = append(p.Pushers, senders...)
}

// 批量执行推送任务
func (p *PushList) Exec(om cm.OutMessage) {
	w := sync.WaitGroup{}
	for _, pp := range p.Pushers {
		w.Add(1)
		go func(om cm.OutMessage, pp sender.Pusher) {
			defer w.Done()
			err := pp.Push(om)
			global.Logger.Error("it occurs error when push messgae", zap.Error(err))
		}(om, pp)
	}
	w.Wait()
}
