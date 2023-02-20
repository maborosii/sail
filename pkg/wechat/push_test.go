package wechat

import (
	"fmt"
	"os"
	"testing"
)

func TestQYWeChatPusher_Push(t *testing.T) {
	type args struct {
		m *QYWeChatMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "qy-chat message type is markdown",
			args: args{
				NewQYWeChatMessage(WithQYWeChatMessageType(MsgTypeMarkdown), WithQYWeChatMessageContentOfMarkDown(struct {
					Content string "json:\"content\""
				}{
					Content: "# test title level 1\n> - test comment\n> - test maintext"})),
			},
			wantErr: false,
		},
	}
	accessToken, _ := os.ReadFile(".access_token")
	d := &QYWeChatConfig{
		AccessToken: string(accessToken),
		Domain:      "https://qyapi.weixin.qq.com/cgi-bin/webhook/send",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewQYWeChatPusher(d)
			fmt.Printf("%+v", p)
			if err := p.Push(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("DingTalkPusher.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
