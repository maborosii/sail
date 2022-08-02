package service

import (
	"encoding/json"
	"fmt"
	"sail/internal/model"
	"testing"
)

func TestService_HarborUploadChart(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test for json descode"},
	}
	huc := &model.HarborUploadRequest{}
	jsonForHuc := `{
		"type": "UPLOAD_CHART",
		"occur_at": 1655959720,
		"operator": "admin",
		"event_data": {
			"resources": [{
				"tag": "0.1.0",
				"resource_url": "http://harbor.domain.com:5002/chartrepo/chart-dg/charts/ale-task-job-executor-supervision-0.1.0.tgz"
			}],
			"repository": {
				"name": "ale-task-job-executor-supervision",
				"namespace": "chart-dg",
				"repo_full_name": "chart-dg/ale-task-job-executor-supervision",
				"repo_type": "public"
			}
		}
	}`
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(jsonForHuc), huc)
			fmt.Println(err)
		})
	}
}
