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

type Pusher struct {
	access_token string
	secret       string

	Client *http.Client
}

func NewPusher(access_token string, secret string) *Pusher {
	return &Pusher{
		access_token: access_token,
		secret:       secret,
	}
}

type PushMessageType string

const (
	MSG_TYPE_TEXT PushMessageType = "text"
)

type PushError struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func (perr *PushError) Error() string {
	return fmt.Sprintf("DingTalk API error code:%d msg:%q", perr.Code, perr.Message)
}

type PushMessage struct {
	access_token string
	secret       string

	Type PushMessageType `json:"msgtype"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewPushMessage(msgtype PushMessageType) PushMessage {
	m := PushMessage{Type: msgtype}
	return m
}

func (m PushMessage) SetText(content string) PushMessage {
	m.Text.Content = content
	return m
}
func (m PushMessage) SetAccessToken(token string) PushMessage {
	m.access_token = token
	return m
}
func (m PushMessage) SetSecret(secret string) PushMessage {
	m.secret = secret
	return m
}

func (pusher *Pusher) Push(m PushMessage) error {
	client := pusher.Client
	if client == nil {
		client = http.DefaultClient
	}

	access_token := m.access_token
	if access_token == "" {
		access_token = pusher.access_token
	}
	secret := m.secret
	if secret == "" {
		secret = pusher.secret
	}

	sign, timestamp := tosign(secret)

	u, err := url.Parse(`https://oapi.dingtalk.com/robot/send`)
	if err != nil {
		panic(err)
	}

	v := url.Values{}
	v.Set("access_token", access_token)
	if sign != "" {
		v.Set("sign", sign)
		v.Set("timestamp", timestamp)
	}
	u.RawQuery = v.Encode()

	buf := bytes.NewBuffer(nil)
	err = json.NewEncoder(buf).Encode(&m)
	if err != nil {
		return err
	}
	fmt.Println(u.String())
	resp, err := client.Post(u.String(), "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	perr := &PushError{}
	json.NewDecoder(resp.Body).Decode(perr)
	if perr.Code != 0 {
		return perr
	}
	return nil
}

func (pusher *Pusher) PushText(text string) error {
	return pusher.Push(NewPushMessage(MSG_TYPE_TEXT).SetText(text))
}

func tosign(secret string) (sign string, timestamp string) {
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
