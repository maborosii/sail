// 定义全局变量
package global

import (
	"time"

	"go.uber.org/zap"
)

var (
	// MonitorSetting *setting.MonitorConfig
	Logger     *zap.Logger
	ConfigPath string
	RateLimit  time.Duration
	BucketSize int
)
