package starters

import (
	"fmt"
	"log"
	"sail/global"
	"sail/infra"
	sd "sail/internal/sender"
	"sail/pkg/setting"
	wc "sail/pkg/wechat"

	"github.com/mitchellh/mapstructure"
)

type QYWeChatStarter struct {
	infra.BaseStarter
}

func (q *QYWeChatStarter) Setup(conf *setting.Config) {
	q.setupPusher(conf)
	q.setupTemplate(conf)
}

func (q *QYWeChatStarter) setupPusher(conf *setting.Config) {
	log.Println("init qy-wechat pusher...")
	dtc := &wc.QYWeChatConfig{}
	for k, v := range conf.Senders {
		if k == "qywechat" {
			err := mapstructure.Decode(v, dtc)
			if err != nil {
				panic(fmt.Sprintf("mapstructure qy-wechat config occur %s", err))
			}
		}
	}
	global.PusherOfQYWeChat = wc.NewQYWeChatPusher(dtc)
	sd.PusherList.RegisterPusher(global.PusherOfQYWeChat)
}

func (q *QYWeChatStarter) setupTemplate(conf *setting.Config) {
	for tt, j := range conf.Template {
		if tt == "qywechat" {
			for tn, jj := range j {
				switch tn {
				case "harbor_upload_chart":
					err := mapstructure.Decode(jj, global.TemplateQYWeChatHarborUploadChart)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for qy-wechat occur %s", err))
					}
				case "harbor_replication_image":
					err := mapstructure.Decode(jj, global.TemplateQYWeChatHarborReplicationImage)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for qy-wechat harbor_replication_image occur %s", err))
					}
				case "harbor_replication_chart":
					err := mapstructure.Decode(jj, global.TemplateQYWeChatHarborReplicationChart)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for qy-wechat harbor_replication_chart occur %s", err))
					}
				case "argocd_sync":
					err := mapstructure.Decode(jj, global.TemplateQYWeChatArgocdSync)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for qy-wechat occur %s", err))
					}
				}
			}
		}
	}
}
