package service

import (
	"sail/global"
	"sail/internal/model"
	sd "sail/internal/sender"
	dts "sail/internal/sender/dingtalk"
	wcs "sail/internal/sender/wechat"
	dt "sail/pkg/dingtalk"
	q "sail/pkg/queue"
	wc "sail/pkg/wechat"

	"go.uber.org/zap"
)

func (s *Service) ArgocdNotify(req *model.ArgocdNotifyRequest) error {
	_, err := req.ArgocdMsgType()
	if err != nil {
		global.Logger.Error("argocd request's type is not support",
			zap.Error(err),
			zap.String("request type", req.Type))
		return err
	}

	// render + pusher
	var hrc = &dt.DingTalkRender{
		Template: global.TemplateDingTalkArgocdSync,
		Render:   dts.Render,
	}
	got, err := hrc.Rend(req, hrc.Template)
	if err != nil {
		global.Logger.Error("rend argocd message occurred err", zap.Error(err))
		return err
	}

	// push job
	j := q.NewJob(func() error {
		sd.PusherList.Exec(got)
		return nil
	})
	go func(j *q.Job) {
		global.FlowControl.CommitJob(j)
		global.Logger.Debug("commit argocd job to job queue success", zap.String("job_uuid", j.UUID))
		j.WaitDone()
		global.Logger.Debug("job has completed", zap.String("job_uuid", j.UUID))
	}(j)
	return nil
}

// TODO: combine to ArgocdNotify
func (s *Service) ArgocdNotify_WeChat(req *model.ArgocdNotifyRequest) error {
	_, err := req.ArgocdMsgType()
	if err != nil {
		global.Logger.Error("argocd request's type is not support",
			zap.Error(err),
			zap.String("request type", req.Type))
		return err
	}

	// render + pusher
	var hrc = &wc.QYWeChatRender{
		Template: global.TemplateQYWeChatArgocdSync,
		Render:   wcs.Render,
	}
	got, err := hrc.Rend(req, hrc.Template)
	if err != nil {
		global.Logger.Error("rend argocd message occurred err", zap.Error(err))
		return err
	}

	// push job
	j := q.NewJob(func() error {
		sd.PusherList.Exec(got)
		return nil
	})
	go func(j *q.Job) {
		global.FlowControl.CommitJob(j)
		global.Logger.Debug("commit argocd job to job queue success", zap.String("job_uuid", j.UUID))
		j.WaitDone()
		global.Logger.Debug("job has completed", zap.String("job_uuid", j.UUID))
	}(j)
	return nil
}
