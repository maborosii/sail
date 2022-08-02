package starters

import (
	"fmt"
	"sail/global"
	"sail/infra"
	q "sail/pkg/queue"
	"sail/pkg/setting"
)

type FlowControlStarter struct {
	infra.BaseStarter
}

func (f *FlowControlStarter) Setup(conf *setting.Config) {
	fmt.Println("init Flow control ..")
	global.FlowControl = q.NewFlowControl()
}
