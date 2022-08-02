// 定义全局变量
package global

import (
	dt "sail/pkg/dingtalk"
	q "sail/pkg/queue"

	"go.uber.org/zap"
)

var (
	ConfigPath string = "configs"
)
var (
	Logger      *zap.Logger
	FlowControl *q.FlowControl
)
var (
	PusherOfDingtalk *dt.DingTalkPusher
)

// var (
// 	TemplateDingTalkHarborReplicationImage *dt.DingTalkMessageTemplate
// 	TemplateDingTalkHarborReplicationChart *dt.DingTalkMessageTemplate
// 	TemplateDingTalkHarborUploadChart      *dt.DingTalkMessageTemplate
// 	TemplateDingTalkArgocdSync             *dt.DingTalkMessageTemplate
// )

var (
	TemplateDingTalkHarborReplicationImage = new(dt.DingTalkMessageTemplate)
	TemplateDingTalkHarborReplicationChart = new(dt.DingTalkMessageTemplate)
	TemplateDingTalkHarborUploadChart      = new(dt.DingTalkMessageTemplate)
	TemplateDingTalkArgocdSync             = new(dt.DingTalkMessageTemplate)
)
