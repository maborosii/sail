package service

import (
	"fmt"
	"sail/global"
	"sail/pkg/util"

	"go.uber.org/zap"
)

type ArgocdNotifyRequest struct {
	*CommonRequest
	EventData ArgocdEventData `json:"event_data"`
}

type ArgocdEventData struct {
	Source       string `json:"source"`
	AppName      string `json:"app_name"`
	SyncStatus   string `json:"sync_status"`
	HealthStatus string `json:"health_status"`
}

func (a *ArgocdNotifyRequest) Spread(keys ...string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	fmt.Printf("%+v\n", a)
	m, err := util.SpreadToMap(*a, "json")
	if err != nil {
		global.Logger.Error("argocd notify request convert to map failed", zap.Error(err))
		return nil, err
	}
	for _, key := range keys {
		out[key] = m[key]
	}
	return out, nil
}

func (s *Service) ArgocdNotify(req *ArgocdNotifyRequest) error { return nil }
