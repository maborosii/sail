package starters

import (
	"log"
	"sail/infra"
	"sail/internal/receiver"
	"sail/pkg/setting"
)

type RecvStarter struct {
	infra.BaseStarter
}

func (r *RecvStarter) Setup(conf *setting.Config) {
	log.Println("init receiver ..")
	// fmt.Println("init receiver ..")
	recv := conf.Receiver
	receiver.Receiver(recv.Port)
}
