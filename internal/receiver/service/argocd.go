package service

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

func (s *Service) ArgocdNotify(req *ArgocdNotifyRequest) error { return nil }
