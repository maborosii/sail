package v1

import (
	"github.com/gin-gonic/gin"
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
func (h Harbor) NotifyChartUpload(c *gin.Context) {}
