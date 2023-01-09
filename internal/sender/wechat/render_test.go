package wechat

import (
	"fmt"
	"os"
	cm "sail/common"
	"sail/internal/model"
	"sail/internal/sender"
	wc "sail/pkg/wechat"

	"testing"
)

func TestRender(t *testing.T) {
	type args struct {
		n           cm.InputMessage
		templateStr *wc.QYWeChatMessageTemplate
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test argocd sync",
			args: args{
				n: &model.ArgocdNotifyRequest{
					CommonRequest: &model.CommonRequest{
						Type:     "argocd",
						Operator: "default",
						OccurAt:  1655959720,
					},
					EventData: model.ArgocdEventData{
						Source:       "本地集群",
						AppName:      "dev-demo",
						SyncStatus:   "Succeed",
						HealthStatus: "Healthy",
					},
				},
				templateStr: &wc.QYWeChatMessageTemplate{
					Content: "### ArgoCD 应用同步状态 -- {{ .City }}\n> - 名称: {{ .AppName }}\n> - 同步状态: {{ .SyncStatus }}\n> - 健康状态: {{ .HealthStatus }}\n> - 时间: {{ .OccurTime }}\n",
				},
			},
		},
	}
	accessToken, _ := os.ReadFile("../../../pkg/wechat/.access_token")
	d := &wc.QYWeChatConfig{
		AccessToken: string(accessToken),
		Domain:      "https://qyapi.weixin.qq.com/cgi-bin/webhook/send",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := Render(tt.args.n, tt.args.templateStr)
			fmt.Printf("%+v\n", m)
			if err := sender.PushMessage(wc.NewQYWeChatPusher(d), m); err != nil {
				fmt.Println(err)
			}
		})
	}
}
