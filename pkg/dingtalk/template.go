package dingtalk

import cm "sail/common"

// pre asset
var a cm.Render
var _ = a.(*DingTalkRender)

type DingTalkMessageTemplate struct {
	MsgType string `mapstructure:"type"`
	Title   string `mapstructure:"title"`
	Content string `mapstructure:"content"`
}

func (dtt *DingTalkMessageTemplate) GetSentence() []string {
	return []string{dtt.Title, dtt.Content}
}

// 渲染器
type DingTalkRender struct {
	Template *DingTalkMessageTemplate
	Render   func(n cm.InputMessage, omt cm.OutMessageTemplate) (cm.OutMessage, error)
}

func (dtr *DingTalkRender) Rend(n cm.InputMessage, omt cm.OutMessageTemplate) (cm.OutMessage, error) {
	return dtr.Render(n, omt)
}
