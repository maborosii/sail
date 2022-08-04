package dingtalk

import (
	"bytes"
	cm "sail/common"
	"sail/global"
	"sail/internal/model"
	dt "sail/pkg/dingtalk"
	"sail/pkg/errcode"
	"text/template"
	"time"

	"go.uber.org/zap"
)

/*
	通过模板文件渲染消息体
*/
func Render(n cm.InputMessage, omt cm.OutMessageTemplate) (cm.OutMessage, error) {
	var expandRequest map[string]interface{}
	var err error

	switch v := n.(type) {
	case *model.ArgocdNotifyRequest:
		expandRequest, err = n.Spread("mapstructure", "city", "app_name", "sync_status", "health_status", "occur_at")
		if err != nil {
			global.Logger.Error("spread argocd notify request to map occur error", zap.Error(err))
			return nil, err
		}
		a := struct {
			City         string
			AppName      string
			SyncStatus   string
			HealthStatus string
			OccurTime    string
			Color        string
		}{
			City:         expandRequest["city"].(string),
			AppName:      expandRequest["app_name"].(string),
			SyncStatus:   expandRequest["sync_status"].(string),
			HealthStatus: expandRequest["health_status"].(string),
			OccurTime:    (time.Unix(int64(expandRequest["occur_at"].(int)), 0)).Format("2006-01-02 15:04:05"),
			Color:        colorMessage(expandRequest["sync_status"].(string)),
		}
		templateStrSlice := omt.GetSentence()
		outputStrSlice, err := rendSlice(templateStrSlice, a)
		if err != nil {
			panic(err)
		}
		m := dt.NewDingTalkMessage(dt.WithDingTalkMessageType(dt.MsgTypeMarkdown), dt.WithDingTalkMessageContentOfMarkDown(struct {
			Title string "json:\"title\""
			Text  string "json:\"text\""
		}{Title: outputStrSlice[0], Text: outputStrSlice[1]}))
		// fmt.Printf("%+v", *m)
		return m, nil

	case *model.HarborReplicationRequest:
		h, ok := n.(*model.HarborReplicationRequest)
		if !ok {
			global.Logger.Error("render harbor replication request asset failed")
		}
		expandRequest, err := n.Spread("mapstructure", "job_status", "resource_type", "dest_domain", "project", "occur_at")
		if err != nil {
			global.Logger.Error("spread harbor replication request to map occur error", zap.Error(err))
			return nil, err
		}
		// 补充字段
		expandRequest["resource_type"] = h.GetResourceType()
		expandRequest["app_name"] = h.GetResourceName()

		a := struct {
			City         string
			AppName      string
			Project      string
			ResourceType string
			JobStatus    string
			OccurTime    string
			Color        string
		}{
			City:         formatDomainCity(expandRequest["dest_domain"].(string), cityList),
			AppName:      expandRequest["app_name"].(string),
			Project:      expandRequest["project"].(string),
			ResourceType: expandRequest["resource_type"].(string),
			JobStatus:    expandRequest["job_status"].(string),
			OccurTime:    (time.Unix(int64(expandRequest["occur_at"].(int)), 0)).Format("2006-01-02 15:04:05"),
			Color:        colorMessage(expandRequest["job_status"].(string)),
		}
		templateStrSlice := omt.GetSentence()
		outputStrSlice, err := rendSlice(templateStrSlice, a)
		if err != nil {
			panic(err)
		}
		m := dt.NewDingTalkMessage(dt.WithDingTalkMessageType(dt.MsgTypeMarkdown), dt.WithDingTalkMessageContentOfMarkDown(struct {
			Title string "json:\"title\""
			Text  string "json:\"text\""
		}{Title: outputStrSlice[0], Text: outputStrSlice[1]}))
		return m, nil

	case *model.HarborUploadRequest:
		uploadChartStatus := "Success"
		expandRequest, err := n.Spread("mapstructure", "app_name", "project", "occur_at")
		if err != nil {
			global.Logger.Error("spread harbor upload request to map occur error", zap.Error(err))
			return nil, err
		}
		a := struct {
			City      string
			AppName   string
			Project   string
			JobStatus string
			OccurTime string
			Color     string
		}{
			City:      formatProjectCity(expandRequest["project"].(string), cityList),
			AppName:   expandRequest["app_name"].(string),
			Project:   expandRequest["project"].(string),
			JobStatus: uploadChartStatus,
			OccurTime: (time.Unix(int64(expandRequest["occur_at"].(int)), 0)).Format("2006-01-02 15:04:05"),
			Color:     colorMessage(uploadChartStatus),
		}
		templateStrSlice := omt.GetSentence()
		outputStrSlice, err := rendSlice(templateStrSlice, a)
		if err != nil {
			panic(err)
		}
		m := dt.NewDingTalkMessage(dt.WithDingTalkMessageType(dt.MsgTypeMarkdown), dt.WithDingTalkMessageContentOfMarkDown(struct {
			Title string "json:\"title\""
			Text  string "json:\"text\""
		}{Title: outputStrSlice[0], Text: outputStrSlice[1]}))
		// fmt.Printf("%+v", *m)
		return m, nil
	default:
		global.Logger.Error("template render occur err, cannot found request type", zap.Any("req_type", v), zap.Error(errcode.RequestTypeNotSupport))
		return nil, errcode.RequestTypeNotSupport
	}
	// return &dt.DingTalkMessage{}, nil
}

func rendSlice(templateStrSlice []string, input interface{}) ([]string, error) {
	var buf bytes.Buffer
	output := make([]string, 0, 3)

	for _, str := range templateStrSlice {
		tmpl, err := template.New("-").Parse(str)
		if err != nil {
			return []string{}, err
		}
		err = tmpl.Execute(&buf, input)
		if err != nil {
			return []string{}, err
		}
		output = append(output, buf.String())
		buf.Reset()
	}
	return output, nil
}
