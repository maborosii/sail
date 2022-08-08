package main

import (
	"fmt"
	"log"
	"sail/cmd"
	"sail/global"
	"sail/infra"
	"sail/infra/starters"

	// _ "sail/internal/sender"
	"sail/pkg/setting"
)

func main() {
	// 获取配置文件
	var err error
	err = cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

	conf, err := setting.NewSetting(global.ConfigPath)
	if err != nil {
		panic(fmt.Sprintf("get config from %s, occurred err; %s", global.ConfigPath, err))
	}
	confStruct := &setting.Config{}
	if err = conf.ReadConfig(confStruct); err != nil {
		panic(err)
	}
	app := infra.NewBootApplication(confStruct)
	app.Run()
	// select {}
}

// type MyHandler struct {
// 	flowControl *q.FlowControl
// }

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("receieve http request")
// 	job := &q.Job{
// 		UUID:     uuid.NewString(),
// 		DoneChan: make(chan struct{}, 1),
// 		HandleJob: func() error {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.Write([]byte("Hello World"))
// 			time.Sleep(5 * time.Second)
// 			return nil
// 		},
// 	}

// 	h.flowControl.CommitJob(job)
// 	fmt.Println("commit job to job queue success")
// 	job.WaitDone()
// }
func init() {
	// 配置初始化
	infra.Register(&starters.LogStarter{}, &starters.FlowControlStarter{}, &starters.DingTalkStarter{})

	// 启动服务
	infra.Register(&starters.RecvStarter{})
}
