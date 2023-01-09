// 定义全局变量
package global

import (
	dt "sail/pkg/dingtalk"
	q "sail/pkg/queue"
	wc "sail/pkg/wechat"

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
	PusherOfQYWeChat *wc.QYWeChatPusher
)

var (
	TemplateDingTalkHarborReplicationImage = new(dt.DingTalkMessageTemplate)
	TemplateDingTalkHarborReplicationChart = new(dt.DingTalkMessageTemplate)
	TemplateDingTalkHarborUploadChart      = new(dt.DingTalkMessageTemplate)
	TemplateDingTalkArgocdSync             = new(dt.DingTalkMessageTemplate)

	TemplateQYWeChatHarborReplicationImage = new(wc.QYWeChatMessageTemplate)
	TemplateQYWeChatHarborReplicationChart = new(wc.QYWeChatMessageTemplate)
	TemplateQYWeChatHarborUploadChart      = new(wc.QYWeChatMessageTemplate)
	TemplateQYWeChatArgocdSync             = new(wc.QYWeChatMessageTemplate)
)
