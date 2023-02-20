package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	cm "sail/common"
	"time"
)

type QYWeChatConfig struct {
	AccessToken string `mapstructure:"access_token"`
	Domain      string `mapstructure:"domain"`
}

// qy WeChat 消息推送器
type QYWeChatPusher struct {
	accessToken string
	domain      string

	Client *http.Client
}

func NewQYWeChatPusher(qwc *QYWeChatConfig) *QYWeChatPusher {
	return &QYWeChatPusher{
		accessToken: qwc.AccessToken,
		domain:      qwc.Domain,
	}
}

// 对 url 拼接
func (p *QYWeChatPusher) completeURL(mainDomain string) (*url.URL, error) {
	u, err := url.Parse(mainDomain)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Set("key", p.accessToken)
	u.RawQuery = v.Encode()
	return u, nil
}

// 消息推送
func (p *QYWeChatPusher) Push(m cm.OutMessage) error {
	time.Sleep(5 * time.Second)
	// asset
	mm, ok := m.(*QYWeChatMessage)
	if !ok {
		return fmt.Errorf("qy-wechat message asset failed")
	}

	if p.Client == nil {
		p.Client = http.DefaultClient
	}

	u, err := p.completeURL(p.domain)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)
	err = json.NewEncoder(buf).Encode(mm)
	if err != nil {
		return err
	}
	resp, err := p.Client.Post(u.String(), "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	qyWeChatError := &qyWeChatError{}

	if err = json.NewDecoder(resp.Body).Decode(qyWeChatError); err != nil {
		return err
	}
	if qyWeChatError.Code != 0 {
		return qyWeChatError
	}
	return nil
}

func (p *QYWeChatPusher) Type() string {
	return "qy-wechat"
}

// qy-wechat 推送返回消息体
type qyWeChatError struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func (d *qyWeChatError) Error() string {
	return fmt.Sprintf("qy-wechat API error code:%d msg:%q", d.Code, d.Message)
}

// 定义 qy-wechat post 消息体
type QYWeChatPostMessageType string

const (
	MsgTypeText     QYWeChatPostMessageType = "text"
	MsgTypeMarkdown QYWeChatPostMessageType = "markdown"
)

type QYWeChatMessage struct {
	Type QYWeChatPostMessageType `json:"msgtype"`
	Text struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
	MarkDown struct {
		Content string `json:"content"`
	} `json:"markdown,omitempty"`
}

// option mode for init qy-wechat post message body
type OptionOfQYWeChatMessage func(*QYWeChatMessage)

func WithQYWeChatMessageType(t QYWeChatPostMessageType) OptionOfQYWeChatMessage {
	return func(d *QYWeChatMessage) {
		d.Type = t
	}
}
func WithQYWeChatMessageContentOfText(tt struct {
	Content string `json:"content"`
}) OptionOfQYWeChatMessage {
	return func(d *QYWeChatMessage) {
		d.Text = tt
	}
}
func WithQYWeChatMessageContentOfMarkDown(tm struct {
	Content string `json:"content"`
}) OptionOfQYWeChatMessage {
	return func(d *QYWeChatMessage) {
		d.MarkDown = tm
	}
}

func NewQYWeChatMessage(opts ...OptionOfQYWeChatMessage) *QYWeChatMessage {
	m := &QYWeChatMessage{}
	for _, op := range opts {
		op(m)
	}
	return m
}

func (m *QYWeChatMessage) RealText() {
}
