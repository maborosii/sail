package main

import (
	"fmt"
	"sail/global"
	"sail/infra"
	"sail/infra/starters"

	// _ "sail/internal/sender"
	"sail/pkg/setting"
)

func main() {
	// flowControl := q.NewFlowControl()
	// myHandler := MyHandler{
	// 	flowControl: flowControl,
	// }
	// http.Handle("/", &myHandler)

	// http.ListenAndServe(":8080", nil)
	configPath := global.ConfigPath
	conf, err := setting.NewSetting(configPath)
	if err != nil {
		panic(fmt.Sprintf("get config from %s, occur err; %s", configPath, err))
	}
	confStruct := &setting.Config{}
	if err = conf.ReadConfig(confStruct); err != nil {
		panic(err)
	}
	app := infra.NewBootApplication(confStruct)
	// fmt.Printf("%+v", *global.PusherOfDingtalk)
	app.Run()
	// select {}
}

// type MyHandler struct {
// 	flowControl *q.FlowControl
// }

// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("recieve http request")
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
