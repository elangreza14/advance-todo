package config

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	IRedisKeyType string

	IRedisKey struct {
		Types    IRedisKeyType
		Key      string
		Data     interface{}
		Duration time.Duration
		IsNull   bool
	}

	IRedis interface {
		apply(*Configuration) error
		formatKey(req IRedisKey) string
		Set(ctx context.Context, key *IRedisKey) error
		Get(ctx context.Context, key *IRedisKey) error
	}

	iRedis struct {
		rdb *redis.Client
	}
)

const (
	TokenKey IRedisKeyType = "token"
)

func newRedis() IRedis {
	return &iRedis{}
}

func WithRedis() Option {
	return newRedis()
}

func (i *iRedis) apply(conf *Configuration) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Env.REDIS_HOSTNAME, conf.Env.REDIS_PORT),
		Password: conf.Env.REDIS_PASS,
		DB:       conf.Env.REDIS_DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return err
	}

	i.rdb = rdb
	conf.Cache = i
	return nil
}

func (i *iRedis) Set(ctx context.Context, req *IRedisKey) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := i.rdb.Set(ctx, i.formatKey(*req), data, req.Duration).Err(); err != nil {
		return err
	}
	return nil
}

func (i *iRedis) Get(ctx context.Context, req *IRedisKey) error {
	val, err := i.rdb.Get(ctx, i.formatKey(*req)).Result()
	if err == redis.Nil {
		req.Data = nil
		req.IsNull = true
		return nil
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(val), &req)
	if err != nil {
		return err
	}
	req.IsNull = false

	return nil

}

func (i *iRedis) formatKey(req IRedisKey) string {
	return fmt.Sprintf("%v|%v", req.Types, req.Key)
}
