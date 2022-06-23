package common

type HarborMessage struct {
	Type      string        `json:"type"`
	OccurAt   int           `json:"occur_at"`
	Operator  string        `json:"operator"`
	EventData EventData     `json:"event_data"`
	Status    MessageStatus `json:"status,omitempty"`
}
type EventData struct {
	Replication Replication `json:"replication,omitempty"`
	Repository  Repository  `json:"repository,omitempty"`
	Resources   Resources   `json:"resources,omitempty"`
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

func (h *HarborMessage) messageType() harborMessageType {
	switch h.Type {
	case string(REPLICATION):
		return REPLICATION
	case string(UPLOAD_CHART):
		return UPLOAD_CHART
	default:
		return "unknown message type"
	}
}

// initial harbor message type
func (h *HarborMessage) InitStatus() {
	h.Status = RECEIVED
}

// get harbor message type
func (h *HarborMessage) GetStatus() (MessageStatus, error) {
	// if harbor message status is not set
	if h.Status == NOT_SET {
		return NOT_SET, Err_NOT_SET_MSG_STATUS
	}
	return h.Status, nil
}

// transform harbor message type, example: RECEIVED -> DEPLOYED
func (h *HarborMessage) ConvertStatus(destStatus MessageStatus) {
	if h.Status != destStatus {
		h.Status = destStatus
	}
}

func (h *HarborMessage) GetName() (string, error) {
	if h.messageType() != UPLOAD_CHART {
		return "", Err_NOT_SUPPORT_MSG_TYPE
	}
	return h.EventData.Repository.Name, nil
}

func (h *HarborMessage) GetResource() (string, error) {
	if h.messageType() != UPLOAD_CHART {
		return "", Err_NOT_SUPPORT_MSG_TYPE
	}
	return h.EventData.Repository.RepoFullName, nil
}
