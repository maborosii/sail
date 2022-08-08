package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	cm "sail/common"
	"time"
)

type DingTalkConfig struct {
	// Delay       int    `mapstructure:"delay"`
	AccessToken string `mapstructure:"access_token"`
	Secret      string `mapstructure:"secret"`
	Domain      string `mapstructure:"domain"`
}

// dingtalk 消息推送器
type DingTalkPusher struct {
	accessToken string
	secret      string
	domain      string

	Client *http.Client
}

func NewDingTalkPusher(dtc *DingTalkConfig) *DingTalkPusher {
	return &DingTalkPusher{
		accessToken: dtc.AccessToken,
		secret:      dtc.Secret,
		domain:      dtc.Domain,
	}
}

// 对 url 加签并拼接
func (p *DingTalkPusher) completeURL(mainDomain string) (*url.URL, error) {
	sign, timestamp := toSign(p.secret)

	u, err := url.Parse(mainDomain)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Set("access_token", p.accessToken)
	if sign != "" {
		v.Set("sign", sign)
		v.Set("timestamp", timestamp)
	}
	u.RawQuery = v.Encode()
	return u, nil
}

// 消息推送
func (p *DingTalkPusher) Push(m cm.OutMessage) error {
	// sleep 5 secs
	time.Sleep(5 * time.Second)
	// asset
	mm, ok := m.(*DingTalkMessage)
	if !ok {
		return fmt.Errorf("dingtalk message asset failed")
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

	dingError := &dingTalkError{}

	if err = json.NewDecoder(resp.Body).Decode(dingError); err != nil {
		return err
	}
	if dingError.Code != 0 {
		return dingError
	}
	return nil
}

func (p *DingTalkPusher) Type() string {
	return "dingtalk"
}

// dingtalk 推送返回消息体
type dingTalkError struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func (d *dingTalkError) Error() string {
	return fmt.Sprintf("DingTalk API error code:%d msg:%q", d.Code, d.Message)
}

// 定义 dingtalk post 消息体
type DingTalkPostMessageType string

const (
	MsgTypeText     DingTalkPostMessageType = "text"
	MsgTypeMarkdown DingTalkPostMessageType = "markdown"
)

type DingTalkMessage struct {
	Type DingTalkPostMessageType `json:"msgtype"`
	Text struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
	MarkDown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown,omitempty"`
}

// option mode for init dingtalk post message body
type OptionOfDingTalkMessage func(*DingTalkMessage)

func WithDingTalkMessageType(t DingTalkPostMessageType) OptionOfDingTalkMessage {
	return func(d *DingTalkMessage) {
		d.Type = t
	}
}
func WithDingTalkMessageContentOfText(tt struct {
	Content string `json:"content"`
}) OptionOfDingTalkMessage {
	return func(d *DingTalkMessage) {
		d.Text = tt
	}
}
func WithDingTalkMessageContentOfMarkDown(tm struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}) OptionOfDingTalkMessage {
	return func(d *DingTalkMessage) {
		d.MarkDown = tm
	}
}

func NewDingTalkMessage(opts ...OptionOfDingTalkMessage) *DingTalkMessage {
	m := &DingTalkMessage{}
	for _, op := range opts {
		op(m)
	}
	return m
}

func (m *DingTalkMessage) RealText() {
}

// dingtalk 加签
func toSign(secret string) (sign string, timestamp string) {
	if secret == "" {
		return "", ""
	}
	timestamp = fmt.Sprintf("%d", time.Now().Unix()*1000)
	sign = fmt.Sprintf("%s\n%s", timestamp, secret)
	signData := computeHmacSha256(sign, secret)
	return url.QueryEscape(signData), timestamp
}

func computeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
