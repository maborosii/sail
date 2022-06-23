package routers

import (
	v1 "sail/internal/receiver/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	recv := v1.NewReceiver()
	apiv1 := r.Group("/api/v1")

	{
		apiv1.POST("/harbor/replciation/image", recv.NotifyImageReplication)
		apiv1.POST("/harbor/replciation/chart", recv.NotifyChartReplicaiton)
		apiv1.POST("/harbor/upload/chart", recv.NotifyChartUpload)

	}

	return r
}
