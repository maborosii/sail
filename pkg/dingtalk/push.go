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

// dingtalk 消息推送器
type DingTalkPusher struct {
	access_token string
	secret       string

	Client *http.Client
}

func NewDingTalkPusher(access_token string, secret string) *DingTalkPusher {
	return &DingTalkPusher{
		access_token: access_token,
		secret:       secret,
	}
}

// 对 url 加签并拼接
func (p *DingTalkPusher) completeUrl(mainDomain string) (*url.URL, error) {
	fmt.Println("p.secret:", p.secret)
	sign, timestamp := toSign(p.secret)

	u, err := url.Parse(mainDomain)
	if err != nil {
		panic(err)
	}
	v := url.Values{}
	v.Set("access_token", p.access_token)
	if sign != "" {
		v.Set("sign", sign)
		v.Set("timestamp", timestamp)
	}
	u.RawQuery = v.Encode()
	return u, nil
}

// 消息推送
func (p *DingTalkPusher) Push(m cm.OutMessage) error {
	// func (p *DingTalkPusher) Push(m *DingTalkMessage) error {

	//asset
	mm := m.(*DingTalkMessage)
	if p.Client == nil {
		p.Client = http.DefaultClient
	}

	u, err := p.completeUrl(`https://oapi.dingtalk.com/robot/send`)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)
	// fmt.Printf("%+v\n", mm)
	err = json.NewEncoder(buf).Encode(mm)
	if err != nil {
		return err
	}
	fmt.Println(u.String())
	resp, err := p.Client.Post(u.String(), "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dingError := &dingTalkError{}
	json.NewDecoder(resp.Body).Decode(dingError)
	if dingError.code != 0 {
		return dingError
	}
	return nil
}

// dingtalk 推送返回消息体
type dingTalkError struct {
	code    int    `json:"errcode"`
	message string `json:"errmsg"`
}

func (d *dingTalkError) Error() string {
	return fmt.Sprintf("DingTalk API error code:%d msg:%q", d.code, d.message)
}

// 定义 dingtalk post 消息体
type DingTalkPostMessageType string

const (
	MSG_TYPE_TEXT     DingTalkPostMessageType = "text"
	MSG_TYPE_MARKDOWN DingTalkPostMessageType = "markdown"
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

// func (pusher *DingTalkPusher) PushText(text string) error {
// return pusher.Push(NewDingTalkMessage(MSG_TYPE_TEXT).SetText(text))
// }

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
