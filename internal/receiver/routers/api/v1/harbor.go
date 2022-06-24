package v1

import (
	"sail/global"
	"sail/internal/receiver/service"
	"sail/pkg/errcode"
	resp "sail/pkg/receiver"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Harbor struct{}

func NewHarbor() Harbor {
	return Harbor{}
}

// @Summary harbor镜像复制处理 -- from 源harbor的replication的webhook
// @Produce  json
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/harbor/replication/image [post]
func (h Harbor) NotifyImageReplication(c *gin.Context) {}

// @Summary harbor chart复制处理 -- from 源harbor的replication的webhook
// @Produce  json
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/harbor/replication/chart [post]
func (h Harbor) NotifyChartReplicaiton(c *gin.Context) {}

// @Summary harbor chart上传处理 -- from 当前harbor的upload_chart的webhook
// @Produce  json
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/harbor/upload/chart [post]
func (h Harbor) NotifyChartUpload(c *gin.Context) {
	var err error
	req := service.UploadRequest{}
	response := resp.NewResponse(c)

	err = c.ShouldBindJSON(&req)
	if err != nil {
		global.Logger.Error("request cannot map to struct",
			zap.Error(errcode.BadRequest))
		response.ToErrorResponse(errcode.BadRequest)
		return
	}

	// request validate
	msgType, err := req.MsgType()
	if err != nil {
		global.Logger.Error("upload request's type is not support",
			zap.Error(err),
			zap.String("request type", req.Type))
		response.ToErrorResponse(errcode.RequestTypeNotSupport)
		return
	}
	if msgType != service.UPLOAD_CHART {
		global.Logger.Error("upload request's type is not adaption",
			zap.Error(errcode.RequestTypeNotAdapt),
			zap.String("request type", req.Type))
		response.ToErrorResponse(errcode.RequestTypeNotAdapt)
		return
	}

	// main logic
	srv := service.NewService(c.Request.Context())
	srv.UploadChart(&req)

	global.Logger.Info("request handle successful")
	response.ToResponse(gin.H{})
}
