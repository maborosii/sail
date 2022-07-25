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
	"time"
)

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

func (p *DingTalkPusher) completeUrl(mainDomain string) (*url.URL, error) {
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

func (p *DingTalkPusher) Push(m DingTalkMessage) error {
	if p.Client == nil {
		p.Client = http.DefaultClient
	}

	u, err := p.completeUrl(`https://oapi.dingtalk.com/robot/send`)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)
	err = json.NewEncoder(buf).Encode(&m)
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

type DingTalkPushMessageType string

const (
	MSG_TYPE_TEXT     DingTalkPushMessageType = "text"
	MSG_TYPE_MARKDOWN DingTalkPushMessageType = "markdown"
)

type dingTalkError struct {
	code    int    `json:"errcode"`
	message string `json:"errmsg"`
}

func (d *dingTalkError) Error() string {
	return fmt.Sprintf("DingTalk API error code:%d msg:%q", d.code, d.message)
}

type DingTalkMessage struct {
	Type DingTalkPushMessageType `json:"msgtype"`
	Text struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown,omitempty"`
}

func NewDingTalkMessage(msgtype DingTalkPushMessageType) *DingTalkMessage {
	m := &DingTalkMessage{Type: msgtype}
	return m
}
func (m DingTalkMessage) String() {

}

func (m DingTalkMessage) SetText(content string) DingTalkMessage {
	m.Text.Content = content
	return m
}

func (pusher *DingTalkPusher) PushText(text string) error {
	return pusher.Push(NewDingTalkMessage(MSG_TYPE_TEXT).SetText(text))
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
