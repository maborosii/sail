package service

import (
	"sail/global"
	"sail/pkg/errcode"

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
	EventData UploadEventData `json:"event_data"`
}
type HarborReplicationRequest struct {
	*CommonRequest
	EventData ReplicationEventData `json:"event_data"`
}

type UploadEventData struct {
	Repository Repository `json:"repository,omitempty"`
	Resources  Resources  `json:"resources,omitempty"`
}

type ReplicationEventData struct {
	Replication Replication `json:"replication,omitempty"`
}
type Repository struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	RepoFullName string `json:"repo_full_name"`
	RepoType     string `json:"repo_type"`
}
type Resources struct {
	Resources []Resource `json:"resources"`
}
type Resource struct {
	Tag         string `json:"tag"`
	ResourceURL string `json:"resource_url"`
}

type Replication struct {
	HarborHostname     string               `json:"harbor_hostname"`
	JobStatus          string               `json:"job_status"`
	ArtifactType       string               `json:"artifact_type"`
	AuthenticationType string               `json:"authentication_type"`
	OverrideMode       bool                 `json:"override_mode"`
	TriggerType        string               `json:"trigger_type"`
	ExecutionTimestamp int                  `json:"execution_timestamp"`
	SrcResource        SrcResource          `json:"src_resource"`
	DestResource       DestResource         `json:"dest_resource"`
	SuccessfulArtifact []SuccessfulArtifact `json:"successful_artifact"`
}
type SrcResource struct {
	RegistryType string `json:"registry_type"`
	Endpoint     string `json:"endpoint"`
	Namespace    string `json:"namespace"`
}
type DestResource struct {
	RegistryName string `json:"registry_name"`
	RegistryType string `json:"registry_type"`
	Endpoint     string `json:"endpoint"`
	Namespace    string `json:"namespace"`
}
type SuccessfulArtifact struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	NameTag string `json:"name_tag"`
}

// argocd message define

/*
	service
*/
func (s *Service) HarborUploadChart(req *HarborUploadRequest) error {
	msgType, err := req.HarborMsgType()
	if err != nil {
		global.Logger.Error("upload request's type is not support",
			zap.Error(err),
			zap.String("request type", req.Type))
		return err
	}
	if msgType != UPLOAD_CHART {
		global.Logger.Error("upload request's type is not adaption",
			zap.Error(errcode.RequestTypeNotSupport),
			zap.String("request type", req.Type))
		return errcode.RequestTypeNotSupport
	}
	// TODO
	// 1. 将请求内容简化，并添加状态值
	// 2. 将简化后的内容传送给deployer
	return nil

}

func (s *Service) HarborReplicationChart(req *HarborReplicationRequest) error { return nil }
func (s *Service) HarborReplicationImage(req *HarborReplicationRequest) error { return nil }
