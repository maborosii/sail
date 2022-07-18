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
	Type     string `json:"type"`
	OccurAt  int    `json:"occur_at"`
	Operator string `json:"operator"`
}
