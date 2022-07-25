package service

type ArgocdNotifyRequest struct {
	*CommonRequest
	EventData ArgocdEventData `json:"event_data"`
}

type ArgocdEventData struct {
	Name         string `json:"name"`
	SyncStatus   string `json:"sync_status"`
	HealthStatus string `json:"health_status"`
}

func (s *Service) ArgocdNotify(req *ArgocdNotifyRequest) error { return nil }
