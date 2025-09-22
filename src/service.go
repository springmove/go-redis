package src

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/springmove/sptty"
)

type Service struct {
	sptty.BaseService

	cfg Config

	clients []*ClientContext
}

func (s *Service) ServiceName() string {
	return ServiceGoRedis
}

func (s *Service) Init(app sptty.ISptty) error {
	if err := app.GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	if err := s.initClients(); err != nil {
		return err
	}

	return nil
}

func (s *Service) initClients() error {
	for _, v := range s.cfg.Configs {

		c := ClientContext{
			ctx: context.Background(),

			client: redis.NewClient(&redis.Options{
				Addr:     v.Entry,
				Password: v.Pwd,
				DB:       v.DB,
			}),
		}

		s.clients = append(s.clients, &c)
	}

	return nil
}

func (s *Service) ClientContext(index int) *ClientContext {
	return s.clients[index]
}
