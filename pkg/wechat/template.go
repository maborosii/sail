package wechat

import cm "sail/common"

// pre asset
// var a cm.Render
// var _ = a.(*QYWeChatRender)

type QYWeChatMessageTemplate struct {
	MsgType string `mapstructure:"type"`
	Content string `mapstructure:"content"`
}

func (qmt *QYWeChatMessageTemplate) GetSentence() []string {
	return []string{qmt.Content}
}

// 渲染器
type QYWeChatRender struct {
	Template *QYWeChatMessageTemplate
	Render   func(n cm.InputMessage, omt cm.OutMessageTemplate) (cm.OutMessage, error)
}

func (qwr *QYWeChatRender) Rend(n cm.InputMessage, omt cm.OutMessageTemplate) (cm.OutMessage, error) {
	return qwr.Render(n, omt)
}
