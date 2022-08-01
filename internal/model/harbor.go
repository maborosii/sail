package model

import (
	"sail/global"
	"sail/pkg/errcode"
	"sail/pkg/util"
	"strings"

	"go.uber.org/zap"
)

type harborMessageType string

/*
   request model
*/
// harbor request define
func (c *CommonRequest) HarborMsgType() (harborMessageType, error) {
	switch c.Type {
	case string(REPLICATION):
		return REPLICATION, nil
	case string(UPLOAD_CHART):
		return UPLOAD_CHART, nil
	default:
		return "", errcode.RequestTypeNotSupport
	}
}

type HarborUploadRequest struct {
	*CommonRequest
	EventData UploadEventData `json:"event_data" mapstructure:""`
}
type HarborReplicationRequest struct {
	*CommonRequest
	EventData ReplicationEventData `json:"event_data" mapstructure:""`
}

type UploadEventData struct {
	Repository Repository `json:"repository,omitempty" mapstructure:""`
	Resources  Resources  `json:"resources,omitempty" mapstructure:""`
}

type ReplicationEventData struct {
	Replication Replication `json:"replication,omitempty" mapstructure:""`
}
type Repository struct {
	Name         string `json:"name" mapstructure:"app_name"`
	Namespace    string `json:"namespace" mapstructure:"project"`
	RepoFullName string `json:"repo_full_name" mapstructure:""`
	RepoType     string `json:"repo_type" mapstructure:""`
}
type Resources struct {
	Resources []Resource `json:"resources" mapstructure:""`
}
type Resource struct {
	Tag         string `json:"tag" mapstructure:"chart_tag"`
	ResourceURL string `json:"resource_url" mapstructure:"app_chart_url"`
}

type Replication struct {
	HarborHostname     string               `json:"harbor_hostname" mapstructure:""`
	JobStatus          string               `json:"job_status" mapstructure:"job_status"`
	ArtifactType       string               `json:"artifact_type" mapstructure:"resource_type"`
	AuthenticationType string               `json:"authentication_type" mapstructure:""`
	OverrideMode       bool                 `json:"override_mode" mapstructure:""`
	TriggerType        string               `json:"trigger_type" mapstructure:""`
	ExecutionTimestamp int                  `json:"execution_timestamp" mapstructure:""`
	SrcResource        SrcResource          `json:"src_resource"  mapstructure:""`
	DestResource       DestResource         `json:"dest_resource" mapstructure:"dest_resource"`
	SuccessfulArtifact []SuccessfulArtifact `json:"successful_artifact"  mapstructure:"success_artifact"`
}
type SrcResource struct {
	RegistryType string `json:"registry_type"  mapstructure:""`
	Endpoint     string `json:"endpoint"  mapstructure:""`
	Namespace    string `json:"namespace" mapstructure:""`
}
type DestResource struct {
	RegistryName string `json:"registry_name"  mapstructure:"dest_domain"`
	RegistryType string `json:"registry_type" mapstructure:""`
	Endpoint     string `json:"endpoint" mapstructure:""`
	Namespace    string `json:"namespace" mapstructure:"project"`
}
type SuccessfulArtifact struct {
	Type    string `json:"type"  mapstructure:"success_resource_type"`
	Status  string `json:"status"  mapstructure:"success_job_status"`
	NameTag string `json:"name_tag"  mapstructure:"success_name_tag"`
}

func (h *HarborReplicationRequest) GetResourceType() string {
	// for _, resource := range h.EventData.Replication.SuccessfulArtifact {
	// if resource.Type == "artifact" {
	// return "Docker 镜像"
	// }
	// if resource.Type == "chart" {
	// return "Helm Chart"
	// }
	// }
	// return "Unknown Resources"
	switch h.EventData.Replication.ArtifactType {
	case "artifact":
		return "Docker Image"
	case "chart":
		return "Helm Chart"
	default:
		return "Unknown Resources"
	}
}

func (h *HarborReplicationRequest) GetResourceName() string {
	var successTags []string
	for _, tags := range h.EventData.Replication.SuccessfulArtifact {
		successTags = append(successTags, tags.NameTag)
	}
	return strings.Join(successTags, ";")
}

func (h *HarborUploadRequest) Spread(tagType string, keys ...string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	m, err := util.SpreadToMap(h, tagType)
	if err != nil {
		global.Logger.Error("harbor upload request convert to map failed", zap.Error(err))
		return nil, err
	}
	for _, key := range keys {
		out[key] = m[key]
	}
	return out, nil
}

func (h *HarborReplicationRequest) Spread(tagType string, keys ...string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	m, err := util.SpreadToMap(h, tagType)
	if err != nil {
		global.Logger.Error("harbor replication request convert to map failed", zap.Error(err))
		return nil, err
	}
	for _, key := range keys {
		out[key] = m[key]
	}
	return out, nil
}
