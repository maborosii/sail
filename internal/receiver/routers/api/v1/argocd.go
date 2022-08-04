package v1

import (
	"sail/global"
	"sail/internal/model"
	"sail/internal/receiver/service"
	"sail/pkg/errcode"
	resp "sail/pkg/receiver"
	expire "sail/pkg/validate"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArgoCD struct{}

func NewArgoCD() ArgoCD {
	return ArgoCD{}
}

func (a ArgoCD) NotifyArgocdSyncStatus(c *gin.Context) {
	var err error
	req := model.ArgocdNotifyRequest{}
	response := resp.NewResponse(c)

	err = c.ShouldBindJSON(&req)
	if err != nil {
		global.Logger.Error("request cannot map to struct",
			zap.Error(errcode.BadRequest))
		response.ToErrorResponse(errcode.BadRequest)
		return
	}
	// 检查消息体是否过期
	if _, err := expire.IsExpired(req.OccurAt, 5*time.Minute); err != nil {
		global.Logger.Error("request timestamp is expired", zap.Error(err))
		response.ToErrorResponse(errcode.RequestExpired)
		return
	}

	global.Logger.Debug("request info",
		zap.String("type", req.Type),
		zap.String("appName", req.EventData.AppName),
		zap.String("sync_status", req.EventData.SyncStatus),
		zap.String("health_status", req.EventData.HealthStatus))

	srv := service.NewService(c.Request.Context())
	if err = srv.ArgocdNotify(&req); err != nil {
		global.Logger.Error("error occurred", zap.Error(err))
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	global.Logger.Info("request for NotifyArgocdSyncStatus handle successful")
	response.ToResponse(nil)
}
