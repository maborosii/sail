// 定义全局变量
package global

import (
	"go.uber.org/zap"
)

var (
	// MonitorSetting *setting.MonitorConfig
	Logger     *zap.Logger
	ConfigPath string
)
