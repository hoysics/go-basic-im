package cache

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(endpoint []string, pwd string) (*RedisClient, error) {
	ctx := context.Background()
	opt := &redis.Options{
		Addr:     endpoint[0],
		PoolSize: 10000,
		Password: pwd,
	}
	client := redis.NewClient(opt)
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return &RedisClient{Client: client}, nil
}

func (r *RedisClient) SetBytes(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	cmd := r.Client.Set(ctx, key, value, expiration)
	if cmd == nil {
		return errors.New("redis SetBytes cmd is nil")
	}
	return cmd.Err()
}

func (r *RedisClient) GetBytes(ctx context.Context, key string) ([]byte, error) {
	cmd := r.Client.Get(ctx, key)
	if cmd == nil {
		return nil, errors.New("redis GetBytes cmd is nil")
	}
	data, err := cmd.Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, err
	}
	return data, err
}

func (r *RedisClient) GetUint64(ctx context.Context, key string) (uint64, error) {
	cmd := r.Client.Get(ctx, key)
	if cmd == nil {
		return 0, errors.New("redis GetUint64 cmd is nil")
	}
	return cmd.Uint64()
}

func (r *RedisClient) Del(ctx context.Context, key ...string) error {
	cmd := r.Client.Del(ctx, key...)
	if cmd == nil {
		return errors.New("redis Del cmd is nil")
	}
	return cmd.Err()
}
func (r *RedisClient) SetString(ctx context.Context, key string, value string, expiration time.Duration) error {
	cmd := r.Client.Set(ctx, key, value, expiration)
	if cmd == nil {
		return errors.New("redis SetString cmd is nil")
	}
	return cmd.Err()
}
func (r *RedisClient) GetString(ctx context.Context, key string) (string, error) {
	cmd := r.Client.Get(ctx, key)
	if cmd == nil {
		return "", errors.New("redis GetString cmd is nil")
	}
	return cmd.String(), cmd.Err()
}
func (r *RedisClient) SAdd(ctx context.Context, key string, member interface{}) error {
	cmd := r.Client.SAdd(ctx, key, member)
	if cmd == nil {
		return errors.New("redis SAdd cmd is nil")
	}
	return cmd.Err()
}
func (r *RedisClient) SRem(ctx context.Context, key string, member ...interface{}) error {
	cmd := r.Client.SRem(ctx, key, member)
	if cmd == nil {
		return errors.New("redis SRem cmd is nil")
	}
	return cmd.Err()
}

func (r *RedisClient) SMembersStrSlice(ctx context.Context, key string) ([]string, error) {
	cmd := r.Client.SMembers(ctx, key)
	if cmd == nil {
		return nil, errors.New("redis SMembersStrSlice cmd is nil")
	}
	return cmd.Result()
}

func (r *RedisClient) Incr(ctx context.Context, key string, ttl time.Duration) error {
	_, err := r.Client.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, ttl)
		return nil
	})
	return err
}

func (r *RedisClient) GetKeys(ctx context.Context, key string) ([]string, error) {
	return r.Client.Keys(ctx, key).Result()
}
