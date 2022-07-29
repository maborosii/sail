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
