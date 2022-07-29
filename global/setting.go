// 定义全局变量
package global

import (
	dtt "sail/pkg/dingtalk"
	q "sail/pkg/queue"

	"go.uber.org/zap"
)

var (
	Logger      *zap.Logger
	ConfigPath  string
	BucketSize  int
	FlowControl *q.FlowControl

	TemplateHarborReplicationImage *dtt.DingTalkMessageTemplate
	TemplateHarborReplicationChart *dtt.DingTalkMessageTemplate
	TemplateHarborUploadChart      *dtt.DingTalkMessageTemplate
	TemplateArgocdSync             *dtt.DingTalkMessageTemplate
)
