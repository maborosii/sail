package starters

import (
	"fmt"
	"sail/infra"
	"sail/pkg/logger"
	"sail/pkg/setting"
)

type LogStarter struct {
	infra.BaseStarter
}

func (l *LogStarter) Setup(conf *setting.Config) {
	fmt.Println("init logger ..")
	logger.InitLogger(conf.LogConfig)
}
