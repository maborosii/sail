package v1

import "github.com/gin-gonic/gin"

type ArgoCD struct{}

func NewArgoCD() ArgoCD {
	return ArgoCD{}

}

func (a ArgoCD) NotifySyncStatus(c *gin.Context) {}
