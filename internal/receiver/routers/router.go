package routers

import (
	v1 "sail/internal/receiver/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	harbor := v1.NewHarbor()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/harbor/replication", harbor.NotifyHarborReplication)
		apiv1.POST("/harbor/upload/chart", harbor.NotifyHarborChartUpload)
	}

	argocd := v1.NewArgoCD()
	{
		apiv1.POST("/argocd/notify", argocd.NotifyArgocdSyncStatus)
	}

	return r
}
