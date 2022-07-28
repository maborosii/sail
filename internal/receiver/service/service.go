package service

import "context"

type Service struct {
	ctx context.Context
}

func NewService(ctx context.Context) *Service {
	return &Service{
		ctx: ctx,
	}
}

type CommonRequest struct {
	Type     string `json:"type" mapstructure:"type"`
	OccurAt  int    `json:"occur_at" mapstructure:"occur_at"`
	Operator string `json:"operator" mapstructure:"operator"`
}

type NotifyRequest interface {
	Spread(tagType string, keys ...string) (map[string]interface{}, error)
}
