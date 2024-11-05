package service

import (
	"ApiGateway/internal/clients/redis"
	"context"
	"log/slog"
)

type Auth interface {
	SetToken(ctx context.Context, userID string, token string) error
	GetToken(ctx context.Context, userID string) (string, error)
}

type Service struct {
	Auth
	redis *redis.RedisClient
}

func NewService(logg *slog.Logger, redis *redis.RedisClient) *Service {
	return &Service{
		Auth:  NewAuthService(logg, redis),
		redis: redis,
	}
}
