package src

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	ServiceGoRedis = "go-redis"
)

type IServiceGoRedis interface {
	ClientContext(index int) *ClientContext
}

type ClientContext struct {
	ctx    context.Context
	client *redis.Client
}

func (s *ClientContext) Ctx() *context.Context {
	return &s.ctx
}

func (s *ClientContext) Client() *redis.Client {
	return s.client
}
