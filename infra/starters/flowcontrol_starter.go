package starters

import (
	"sail/global"
	"sail/infra"
	q "sail/pkg/queue"
	"sail/pkg/setting"
)

type FlowControlStarter struct {
	infra.BaseStarter
}

func (f *FlowControlStarter) Setup(conf *setting.Config) {
	global.FlowControl = q.NewFlowControl()
}
