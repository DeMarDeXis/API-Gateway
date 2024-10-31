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

//func (s *AuthService) CreateToken(userId string) (string, error) {
//	token := generateToken() // ваша логика генерации токена
//	err := s.redis.SetToken(context.Background(), userId, token, 24*time.Hour)
//	return token, err
//}
