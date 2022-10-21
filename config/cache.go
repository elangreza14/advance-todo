package config

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	ICacheKeyType string

	ICacheKey struct {
		Types    ICacheKeyType
		Key      string
		Data     interface{}
		Duration time.Duration
		IsNull   bool
	}

	ICache interface {
		apply(*Configuration) error
		formatKey(req ICacheKey) string
		Set(ctx context.Context, key *ICacheKey) error
		Get(ctx context.Context, key *ICacheKey) error
	}

	iCache struct {
		rdb *redis.Client
	}
)

const (
	TokenKey ICacheKeyType = "token"
)

func newCache() ICache {
	return &iCache{}
}

func WithCache() Option {
	return newCache()
}

func (i *iCache) apply(conf *Configuration) error {
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

func (i *iCache) Set(ctx context.Context, req *ICacheKey) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := i.rdb.Set(ctx, i.formatKey(*req), data, req.Duration).Err(); err != nil {
		return err
	}
	return nil
}

func (i *iCache) Get(ctx context.Context, req *ICacheKey) error {
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

func (i *iCache) formatKey(req ICacheKey) string {
	return fmt.Sprintf("%v|%v", req.Types, req.Key)
}
