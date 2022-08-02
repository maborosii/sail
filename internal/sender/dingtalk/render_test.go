package dingtalk

import (
	"fmt"
	"os"
	cm "sail/common"
	"sail/internal/model"
	"sail/internal/sender"
	dt "sail/pkg/dingtalk"

	"testing"
)

func TestRender(t *testing.T) {
	type args struct {
		n           cm.InputMessage
		templateStr *dt.DingTalkMessageTemplate
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test argocd sync",
			args: args{
				n: &model.ArgocdNotifyRequest{
					&model.CommonRequest{
						Type:     "argocd",
						Operator: "default",
						OccurAt:  1655959720,
					},
					model.ArgocdEventData{
						Source:       "清远",
						AppName:      "ale-case-service",
						SyncStatus:   "Succeed",
						HealthStatus: "Healthy",
					},
				},
				templateStr: &dt.DingTalkMessageTemplate{
					Title:   "ArgoCD 应用同步状态 -- {{ .City }}",
					Content: "### ArgoCD 应用同步状态 -- {{ .City }}\n> - 名称: {{ .AppName }}\n> - 同步状态: {{ .SyncStatus }}\n> - 健康状态: {{ .HealthStatus }}\n> - 时间: {{ .OccurTime }}\n",
				},
			},
		},
		{
			name: "test harbor replication image",
			args: args{
				n: &model.HarborReplicationRequest{
					&model.CommonRequest{
						Type:     "REPLICATION",
						Operator: "admin",
						OccurAt:  1655959720,
					},
					model.ReplicationEventData{
						Replication: model.Replication{
							HarborHostname:     "harbor.domain.com",
							JobStatus:          "Success",
							ArtifactType:       "artifact",
							AuthenticationType: "basic",
							OverrideMode:       true,
							TriggerType:        "event_based",
							ExecutionTimestamp: 1655959720,
							SrcResource: model.SrcResource{
								RegistryType: "harbor",
								Endpoint:     "http://harbor.domain.com:5002",
								Namespace:    "app",
							},
							DestResource: model.DestResource{
								RegistryName: "qy.harbor.com",
								RegistryType: "harbor",
								Endpoint:     "http://qy.harbor.com:89",
								Namespace:    "app",
							},
							SuccessfulArtifact: []model.SuccessfulArtifact{
								{
									Type:    "artifact",
									Status:  "Succeed",
									NameTag: "ale-case-service [1 item(s) in total]",
								},
							},
						},
					},
				},
				templateStr: &dt.DingTalkMessageTemplate{
					Title:   "Docker 镜像同步状态 -- {{ .City }}",
					Content: "### Docker 镜像同步状态 -- {{ .City }}\n> - 名称: {{ .AppName }}\n> - 任务状态: {{ .JobStatus }}\n> - 仓库: {{ .Project }}\n> - 时间: {{ .OccurTime }}\n",
				},
			},
		}, {
			name: "test harbor chart image",
			args: args{
				n: &model.HarborReplicationRequest{
					&model.CommonRequest{
						Type:     "REPLICATION",
						Operator: "admin",
						OccurAt:  1655959720,
					},
					model.ReplicationEventData{
						Replication: model.Replication{
							HarborHostname:     "harbor.domain.com",
							JobStatus:          "Failed",
							ArtifactType:       "chart",
							AuthenticationType: "basic",
							OverrideMode:       true,
							TriggerType:        "event_based",
							ExecutionTimestamp: 1655959720,
							SrcResource: model.SrcResource{
								RegistryType: "harbor",
								Endpoint:     "http://harbor.domain.com:5002",
								Namespace:    "chart-qy",
							},
							DestResource: model.DestResource{
								RegistryName: "qy.harbor.com",
								RegistryType: "harbor",
								Endpoint:     "http://qy.harbor.com:89",
								Namespace:    "app",
							},
							SuccessfulArtifact: []model.SuccessfulArtifact{
								{
									Type:    "chart",
									Status:  "Succeed",
									NameTag: "ale-case-service [1 item(s) in total]",
								},
							},
						},
					},
				},
				templateStr: &dt.DingTalkMessageTemplate{
					Title:   "Helm Chart 同步状态 -- {{ .City }}",
					Content: "### Helm Chart 同步状态 -- {{ .City }}\n> - 名称: {{ .AppName }}\n> - 任务状态: {{ .JobStatus }}\n> - 仓库: {{ .Project }}\n> - 时间: {{ .OccurTime }}\n",
				},
			},
		},
		{
			name: "test harbor upload chart",
			args: args{
				n: &model.HarborUploadRequest{
					&model.CommonRequest{
						Type:     "UPLOAD_CHART",
						Operator: "admin",
						OccurAt:  1655959720,
					},
					model.UploadEventData{
						Repository: model.Repository{
							Name:         "ale-case-service",
							Namespace:    "chart-qy",
							RepoFullName: "chart-qy/ale-case-service",
							RepoType:     "public",
						},
						Resources: []model.Resource{
							{
								Tag:         "0.1.0",
								ResourceURL: "http://harbor.domain.com:5002/chartrepo/chart-dg/charts/ale-task-job-executor-supervision-0.1",
							},
						},
					},
				},
				templateStr: &dt.DingTalkMessageTemplate{
					Title:   "Helm Chart 上传状态 -- {{ .City }}",
					Content: "### Helm Chart 上传状态 -- {{ .City }}\n> - 名称: {{ .AppName }}\n> - 任务状态: {{ .JobStatus }}\n> - 仓库: {{ .Project }}\n> - 时间: {{ .OccurTime }}",
				},
			}},
	}
	access_token, _ := os.ReadFile("../../../pkg/dingtalk/.access_token")
	secret, _ := os.ReadFile("../../../pkg/dingtalk/.secret")
	d := &dt.DingTalkConfig{
		AccessToken: string(access_token),
		Secret:      string(secret),
		Domain:      "https://oapi.dingtalk.com/robot/send",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got, err := Render(tt.args.n, tt.args.templateStr)
			m, _ := Render(tt.args.n, tt.args.templateStr)
			fmt.Printf("%+v\n", m)
			sender.PushMessage(dt.NewDingTalkPusher(d), m)

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
