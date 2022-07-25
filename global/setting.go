// 定义全局变量
package global

import (
	q "sail/pkg/queue"

	"go.uber.org/zap"
)

var (
	Logger      *zap.Logger
	ConfigPath  string
	RateLimit   int // seconds
	BucketSize  int
	FlowControl *q.FlowControl
)
