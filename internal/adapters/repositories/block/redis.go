package blkrepo

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"transaction-parser/internal/entity"
)

type RedisRepo struct {
	client *redis.Client
}

func NewRedisRepo(client *redis.Client) *RedisRepo {
	return &RedisRepo{client: client}
}

func (r *RedisRepo) Cache(ctx context.Context, key string, b *entity.Block) error {
	blkJson, err := json.Marshal(b)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, blkJson, 0).Err()
}
