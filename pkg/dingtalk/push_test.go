package dingtalk

import (
	"fmt"
	"os"
	"testing"
)

func TestDingTalkPusher_Push(t *testing.T) {
	type args struct {
		m *DingTalkMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {
		// name: "dingtalk message type is text",
		// args: args{
		// NewDingTalkMessage(WithDingTalkMessageType(MSG_TYPE_TEXT), WithDingTalkMessageContentOfText(struct {
		// Content string "json:\"content\""
		// }{
		// Content: "just test for dingtalk text message type",
		// })),
		// },
		// wantErr: false,
		// }, {
		{
			name: "dingtalk message type is markdown",
			args: args{
				NewDingTalkMessage(WithDingTalkMessageType(MSG_TYPE_MARKDOWN), WithDingTalkMessageContentOfMarkDown(struct {
					Title string "json:\"title\""
					Text  string "json:\"text\""
				}{
					Title: "test markdown",
					Text:  "# test title level 1\n> - test comment\n> - test maintext"})),
			},
			wantErr: false,
		},
	}
	access_token, _ := os.ReadFile(".access_token")
	secret, _ := os.ReadFile(".secret")
	// fmt.Println(string(access_token), string(secret))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewDingTalkPusher(string(access_token), string(secret))
			fmt.Printf("%+v", p)
			if err := p.Push(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("DingTalkPusher.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
