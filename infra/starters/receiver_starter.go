package starters

import (
	"sail/infra"
	"sail/internal/receiver"
	"sail/pkg/setting"
)

type RecvStarter struct {
	infra.BaseStarter
}

func (r *RecvStarter) Setup(conf *setting.Config) {
	recv := conf.Receiver
	receiver.Receiver(recv.Port)
}
