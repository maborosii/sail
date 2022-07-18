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
func (h Harbor) NotifyHarborImageReplication(c *gin.Context) {
	var err error
	req := service.HarborReplicationRequest{}
	response := resp.NewResponse(c)

	err = c.ShouldBindJSON(&req)
	if err != nil {
		global.Logger.Error("request cannot map to struct",
			zap.Error(errcode.BadRequest))
		response.ToErrorResponse(errcode.BadRequest)
		return
	}

	srv := service.NewService(c.Request.Context())
	if err = srv.HarborReplicationImage(&req); err != nil {
		global.Logger.Error("error occured")
		return
	}

	global.Logger.Info("request handle successful")
	response.ToResponse(gin.H{})

}

// @Summary harbor chart复制处理 -- from 源harbor的replication的webhook
// @Produce  json
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/harbor/replication/chart [post]
func (h Harbor) NotifyHarborChartReplication(c *gin.Context) {
	var err error
	req := service.HarborReplicationRequest{}
	response := resp.NewResponse(c)

	err = c.ShouldBindJSON(&req)
	if err != nil {
		global.Logger.Error("request cannot map to struct",
			zap.Error(errcode.BadRequest))
		response.ToErrorResponse(errcode.BadRequest)
		return
	}
	srv := service.NewService(c.Request.Context())
	if err = srv.HarborReplicationChart(&req); err != nil {
		global.Logger.Error("error occured")
		return
	}

	global.Logger.Info("request handle successful")
	response.ToResponse(gin.H{})

}

// @Summary harbor chart上传处理 -- from 当前harbor的upload_chart的webhook
// @Produce  json
// @Success 200 {object} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/harbor/upload/chart [post]
func (h Harbor) NotifyHarborChartUpload(c *gin.Context) {
	var err error
	req := service.HarborUploadRequest{}
	response := resp.NewResponse(c)

	err = c.ShouldBindJSON(&req)
	if err != nil {
		global.Logger.Error("request cannot map to struct",
			zap.Error(errcode.BadRequest))
		response.ToErrorResponse(errcode.BadRequest)
		return
	}

	// main logic
	srv := service.NewService(c.Request.Context())
	if err = srv.HarborUploadChart(&req); err != nil {
		global.Logger.Error("error occured")
		return
	}

	// timeout
	// _, cancel := context.WithTimeout(c, 10*time.Second)
	// defer cancel()

	global.Logger.Info("request handle successful")
	response.ToResponse(gin.H{})
}
