package dingtalk

import (
	"os"
	srv "sail/internal/receiver/service"
	"sail/internal/sender"
	dt "sail/pkg/dingtalk"
	setting "sail/pkg/setting"

	"testing"
)

func TestRender(t *testing.T) {
	type args struct {
		n           srv.NotifyRequest
		templateStr *setting.DingTalkMessageTemplate
	}
	tests := []struct {
		name string
		args args
		// want    cm.OutMessage
		// wantErr bool
	}{
		{
			name: "test",
			args: args{
				n: &srv.ArgocdNotifyRequest{
					&srv.CommonRequest{
						Type:     "argocd",
						Operator: "default",
						OccurAt:  1655959720,
					},
					srv.ArgocdEventData{
						Source:       "清远",
						AppName:      "ale-case-service",
						SyncStatus:   "Succeed",
						HealthStatus: "Healthy",
					},
				},
				templateStr: &setting.DingTalkMessageTemplate{
					Title:   "ArgoCD 应用同步状态 -- {{ .City }}",
					Content: "## ArgoCD 应用同步状态 -- {{ .City }}\n> - 名称: {{ .AppName }}\n> - 同步状态: {{ .SyncStatus }}\n> - 健康状态: {{ .HealthStatus }}\n> - 时间: {{ .OccurTime }}\n",
				},
			},
		},
	}
	access_token, _ := os.ReadFile("../../../pkg/dingtalk/.access_token")
	secret, _ := os.ReadFile("../../../pkg/dingtalk/.secret")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got, err := Render(tt.args.n, tt.args.templateStr)
			m, _ := Render(tt.args.n, tt.args.templateStr)
			sender.PushMessage(dt.NewDingTalkPusher(string(access_token), string(secret)), m)

			// if (err != nil) != tt.wantErr {
			// t.Errorf("Render() error = %v, wantErr %v", err, tt.wantErr)
			// return
			// }
			// if !reflect.DeepEqual(got, tt.want) {
			// t.Errorf("Render() = %v, want %v", got, tt.want)
			// }
		})
	}
}
