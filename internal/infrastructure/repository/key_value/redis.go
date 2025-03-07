package key_value

// Имплементация ключ-значение хранилища Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Repository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepository(client *redis.Client) *Repository {
	ctx := context.Background()
	return &Repository{
		client: client,
		ctx:    ctx,
	}
}

func (r *Repository) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		log.Error(err, "Ошибка получения значения по ключу %s", key)
		return "", err
	}
	return val, nil
}

func (r *Repository) Set(key string, val interface{}, ttl int64) error {
	_, err := r.client.Set(r.ctx, key, val, time.Duration(ttl)*time.Minute).Result()
	if err != nil {
		log.Error(err, "Ошибка установки значения %s по ключу %s", val, key)
		return err
	}
	return nil
}

func (r *Repository) Del(key string) error {
	_, err := r.client.Del(r.ctx, key).Result()
	if err != nil {
		log.Error(err, "Ошибка удаления ключа %s", key)
		return err
	}
	return nil
}
