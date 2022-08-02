package starters

import (
	"sail/infra"
	"sail/pkg/logger"
	"sail/pkg/setting"
)

type LogStarter struct {
	infra.BaseStarter
}

func (l *LogStarter) Setup(conf *setting.Config) {
	logger.InitLogger(conf.LogConfig)
}
