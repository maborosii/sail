package service

import (
	"sail/global"
	"sail/internal/model"
	"sail/internal/sender"
	dts "sail/internal/sender/dingtalk"
	dt "sail/pkg/dingtalk"
	"sail/pkg/errcode"

	"go.uber.org/zap"
)

/*
	service
*/
func (s *Service) HarborUploadChart(req *model.HarborUploadRequest) error {
	msgType, err := req.HarborMsgType()
	if err != nil {
		global.Logger.Error("upload request's type is not support",
			zap.Error(err),
			zap.String("request type", req.Type))
		return err
	}
	if msgType != model.UPLOAD_CHART {
		global.Logger.Error("upload request's type is not adaption",
			zap.Error(errcode.RequestTypeNotSupport),
			zap.String("request type", req.Type))
		return errcode.RequestTypeNotSupport
	}

	// render + pusher
	// TODO
	var hrc = &dt.DingTalkRender{
		Template: global.TemplateHarborUploadChart,
		Render:   dts.Render,
	}
	got, _ := hrc.Rend(req, hrc.Template)
	pusher := dt.NewDingTalkPusher("", "")
	sender.PushMessage(pusher, got)

	// push job
	// j := q.NewJob(func() error { return nil })
	// global.FlowControl.CommitJob(j)
	// fmt.Println("commit job to job queue success")
	// j.WaitDone()
	return nil
}

func (s *Service) HarborReplicationChart(req *model.HarborReplicationRequest) error { return nil }
func (s *Service) HarborReplicationImage(req *model.HarborReplicationRequest) error { return nil }
