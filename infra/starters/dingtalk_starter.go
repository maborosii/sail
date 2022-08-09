package starters

import (
	"fmt"
	"log"
	"sail/global"
	"sail/infra"
	sd "sail/internal/sender"
	dt "sail/pkg/dingtalk"
	"sail/pkg/setting"

	"github.com/mitchellh/mapstructure"
)

type DingTalkStarter struct {
	infra.BaseStarter
}

func (d *DingTalkStarter) Setup(conf *setting.Config) {
	d.setupPusher(conf)
	d.setupTemplate(conf)
}

func (d *DingTalkStarter) setupPusher(conf *setting.Config) {
	// fmt.Println("init dingtalk pusher...")
	log.Println("init dingtalk pusher...")
	dtc := &dt.DingTalkConfig{}
	for k, v := range conf.Senders {
		if k == "dingtalk" {
			err := mapstructure.Decode(v, dtc)
			if err != nil {
				panic(fmt.Sprintf("mapstructure dingtalkconfig occur %s", err))
			}
		}
	}
	// fmt.Printf("%+v\n", *dtc)
	global.PusherOfDingtalk = dt.NewDingTalkPusher(dtc)
	// 注册pusher
	sd.PusherList.RegisterPusher(global.PusherOfDingtalk)
	// fmt.Printf("%+v\n", *global.PusherOfDingtalk)
}

func (d *DingTalkStarter) setupTemplate(conf *setting.Config) {
	for tt, j := range conf.Template {
		// parse template config
		if tt == "dingtalk" {
			for tn, jj := range j {
				switch tn {
				case "harbor_upload_chart":
					err := mapstructure.Decode(jj, global.TemplateDingTalkHarborUploadChart)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for dingtalk_harbor_upload_chart occur %s", err))
					}
				case "harbor_replication_image":
					err := mapstructure.Decode(jj, global.TemplateDingTalkHarborReplicationImage)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for dingtalk_harbor_replication_image occur %s", err))
					}
				case "harbor_replication_chart":
					err := mapstructure.Decode(jj, global.TemplateDingTalkHarborReplicationChart)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for dingtalk_harbor_replication_chart occur %s", err))
					}
				case "argocd_sync":
					err := mapstructure.Decode(jj, global.TemplateDingTalkArgocdSync)
					if err != nil {
						panic(fmt.Sprintf("mapstructure template for dingtalk_argocd_sync occur %s", err))
					}
				}
			}
		}
	}
}
