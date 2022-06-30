package routers

import (
	v1 "sail/internal/receiver/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	harb := v1.NewHarbor()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/harbor/replciation/image", harb.NotifyImageReplication)
		apiv1.POST("/harbor/replciation/chart", harb.NotifyChartReplicaiton)
		apiv1.POST("/harbor/upload/chart", harb.NotifyChartUpload)
	}

	argocd := v1.NewArgoCD()
	{
		apiv1.POST("/argocd/notify", argocd.NotifySyncStatus)
	}

	return r
}
