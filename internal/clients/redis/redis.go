package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

type RedisClient struct {
	client *redis.Client
	log    *slog.Logger
}

func NewRedisClient(addr string, log *slog.Logger) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0, // 1 - 15

		// Connection settings
		//Username:     "",     //Username for ACL(Access Control List)
		//MaxRetries:   3,     //Maximum number of retries
		//DialTimeout:  5 * time.Second,  //Timeout for establishing connections
		//ReadTimeout:  3 * time.Second,  //Timeout for read operations
		//WriteTimeout: 3 * time.Second,  //Timeout for write operations

		// Pool settings
		//PoolSize:     10,    // Maximum number of socket connections
		//MinIdleConns: 2,     // Minimum number of idle connections
		//MaxConnAge:   0,     // Maximum age of connections
		//PoolTimeout:  4 * time.Second,  // Timeout for getting connection from pool

		// TLS settings
		//TLSConfig: nil,      // TLS configuration if needed
	})

	return &RedisClient{client: client}
}

func (r *RedisClient) SetToken(ctx context.Context, userID string, token string, expiration time.Duration) error {
	return r.client.Set(ctx, userID, token, expiration).Err()
}

func (r *RedisClient) GetToken(ctx context.Context, userID string) (string, error) {
	return r.client.Get(ctx, userID).Result()
}
