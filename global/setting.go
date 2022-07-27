// 定义全局变量
package global

import (
	q "sail/pkg/queue"
	"sail/pkg/setting"

	"go.uber.org/zap"
)

var (
	Logger      *zap.Logger
	ConfigPath  string
	BucketSize  int
	FlowControl *q.FlowControl

	TemplateHarborReplicationImage *setting.DingTalkMessageTemplate
	TemplateHarborReplicationChart *setting.DingTalkMessageTemplate
	TemplateHarborUploadChart      *setting.DingTalkMessageTemplate
	TemplateArgocdSync             *setting.DingTalkMessageTemplate
)
