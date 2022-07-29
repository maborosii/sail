package model

type CommonRequest struct {
	Type     string `json:"type" mapstructure:"req_type"`
	OccurAt  int    `json:"occur_at" mapstructure:"occur_at"`
	Operator string `json:"operator" mapstructure:""`
}
