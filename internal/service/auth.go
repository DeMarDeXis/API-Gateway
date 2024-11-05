package service

import (
	"ApiGateway/internal/clients/redis"
	"context"
	"log/slog"
	"time"
)

type AuthService struct {
	logg  *slog.Logger
	redis *redis.RedisClient
}

func NewAuthService(logg *slog.Logger, redis *redis.RedisClient) *AuthService {
	return &AuthService{
		logg:  logg,
		redis: redis,
	}
}

func (s *AuthService) SetToken(ctx context.Context, userID string, token string) error {
	s.logg.Debug("Service_userID: ", userID)
	return s.redis.SetToken(context.Background(), userID, token, 24*time.Hour)
}

func (s *AuthService) GetToken(ctx context.Context, userID string) (string, error) {
	return s.redis.GetToken(context.Background(), userID)
}
