package service

import (
	"sail/global"
	"sail/internal/model"

	sd "sail/internal/sender"
	dts "sail/internal/sender/dingtalk"
	dt "sail/pkg/dingtalk"
	"sail/pkg/errcode"
	q "sail/pkg/queue"

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
	if msgType != model.UPLOADCHART {
		global.Logger.Error("upload request's type is not adaption",
			zap.Error(errcode.RequestTypeNotSupport),
			zap.String("request type", req.Type))
		return errcode.RequestTypeNotSupport
	}

	// render + pusher
	var hrc = &dt.DingTalkRender{
		Template: global.TemplateDingTalkHarborUploadChart,
		Render:   dts.Render,
	}
	got, err := hrc.Rend(req, hrc.Template)
	if err != nil {
		global.Logger.Error("rend upload_chart message occurred err", zap.Error(err))
		return err
	}

	// push job
	j := q.NewJob(func() error {
		sd.PusherList.Exec(got)
		return nil
	})

	go func(j *q.Job) {
		global.FlowControl.CommitJob(j)
		global.Logger.Debug("commit replication job to job queue success", zap.String("job_uuid", j.UUID))
		j.WaitDone()
		global.Logger.Debug("job has completed", zap.String("job_uuid", j.UUID))
	}(j)
	// sd.PusherList.Exec(got)
	// global.PusherOfDingtalk.Push(got)
	return nil
}

func (s *Service) HarborReplication(req *model.HarborReplicationRequest) error {
	msgType, err := req.HarborMsgType()
	if err != nil {
		global.Logger.Error("replication request's type is not support",
			zap.Error(err),
			zap.String("request type", req.Type))
		return err
	}
	if (msgType != model.UPLOADCHART) && (msgType != model.REPLICATION) {
		global.Logger.Error("replication request's type is not adaption",
			zap.Error(errcode.RequestTypeNotSupport),
			zap.String("request type", req.Type))
		return errcode.RequestTypeNotSupport
	}

	// render + pusher
	var replicationTemplate *dt.DingTalkMessageTemplate

	switch req.GetResourceType() {
	case "Helm Chart":
		replicationTemplate = global.TemplateDingTalkHarborReplicationChart
	case "Docker Image":
		replicationTemplate = global.TemplateDingTalkHarborReplicationImage
	default:
		global.Logger.Error("replication request's resource type is not support", zap.String("resource_type", req.EventData.Replication.ArtifactType))
		return errcode.RequestTypeNotSupport
	}

	var hrc = &dt.DingTalkRender{
		Template: replicationTemplate,
		Render:   dts.Render,
	}

	got, err := hrc.Rend(req, hrc.Template)

	if err != nil {
		global.Logger.Error("rend replication message occurred err", zap.Error(err))
		return err
	}

	// push job
	j := q.NewJob(func() error {
		sd.PusherList.Exec(got)
		return nil
	})
	go func(j *q.Job) {
		global.FlowControl.CommitJob(j)
		global.Logger.Debug("commit replication job to job queue success", zap.String("job_uuid", j.UUID))
		j.WaitDone()
		global.Logger.Debug("job has completed", zap.String("job_uuid", j.UUID))
	}(j)
	return nil
}
