// 定义全局变量
package global

import (
	dt "sail/pkg/dingtalk"
	q "sail/pkg/queue"

	"go.uber.org/zap"
)

var (
	Logger      *zap.Logger
	ConfigPath  string
	BucketSize  int
	FlowControl *q.FlowControl

	TemplateHarborReplicationImage *dt.DingTalkMessageTemplate
	TemplateHarborReplicationChart *dt.DingTalkMessageTemplate
	TemplateHarborUploadChart      *dt.DingTalkMessageTemplate
	TemplateArgocdSync             *dt.DingTalkMessageTemplate
)
