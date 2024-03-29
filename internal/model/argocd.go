package model

import (
	"sail/global"
	"sail/pkg/errcode"
	"sail/pkg/util"

	"go.uber.org/zap"
)

func (c *CommonRequest) ArgocdMsgType() (string, error) {
	switch c.Type {
	case "argocd":
		return "argocd", nil
	default:
		return "", errcode.RequestTypeNotSupport
	}
}

type ArgocdNotifyRequest struct {
	*CommonRequest
	EventData ArgocdEventData `json:"event_data" mapstructure:""`
}

type ArgocdEventData struct {
	Source       string `json:"source" mapstructure:"city"`
	AppName      string `json:"app_name" mapstructure:"app_name"`
	SyncStatus   string `json:"sync_status" mapstructure:"sync_status"`
	HealthStatus string `json:"health_status" mapstructure:"health_status"`
}

func (a *ArgocdNotifyRequest) Spread(tagType string, keys ...string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	m, err := util.SpreadToMap(a, tagType)
	if err != nil {
		global.Logger.Error("argocd notify request convert to map failed", zap.Error(err))
		return nil, err
	}
	for _, key := range keys {
		out[key] = m[key]
	}
	return out, nil
}
