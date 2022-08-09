package starters

import (
	"log"
	"sail/infra"
	"sail/pkg/logger"
	"sail/pkg/setting"
)

type LogStarter struct {
	infra.BaseStarter
}

func (l *LogStarter) Setup(conf *setting.Config) {
	log.Println("init logger ..")
	// fmt.Println("init logger ..")
	if err := logger.InitLogger(conf.LogConfig); err != nil {
		panic(err)
	}
}
