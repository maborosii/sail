package common

type MessageStatus int8
type HarborMessage struct {
	Status MessageStatus `json:"status,omitempty"`
}

// initial harbor message type
func (h *HarborMessage) InitStatus() {
	h.Status = RECEIVED
}

// get harbor message type
func (h *HarborMessage) GetStatus() MessageStatus {
	// if harbor message status is not set
	if h.Status == NOTSET {
		return NOTSET
	}
	return h.Status
}

// transform harbor message type, example: RECEIVED -> DEPLOYED
func (h *HarborMessage) ConvertStatus(destStatus MessageStatus) {
	if h.Status != destStatus {
		h.Status = destStatus
	}
}

func (h *HarborMessage) GetName() (string, error) {
	// TODO
	return "", nil
}

func (h *HarborMessage) GetResource() (string, error) {
	// TODO
	return "", nil
}
