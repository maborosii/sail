package common

import "errors"

type MessageStatus int8
type harborMessageType string

const (
	NOT_SET MessageStatus = iota
	RECEIVED
	DEPLOYED
)

const (
	REPLICATION  harborMessageType = "REPLICATION"
	UPLOAD_CHART harborMessageType = "UPLOAD_CHART"
)

var (
	Err_NOT_SUPPORT_MSG_TYPE = errors.New("not support message type")
	Err_NOT_SET_MSG_STATUS   = errors.New("message type not set")
)
