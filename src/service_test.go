package src

import (
	"fmt"
	"testing"
)

func initSrv() IServiceGoRedis {
	s := Service{
		cfg: Config{
			Configs: []RedisEntry{
				{
					Entry: "",
					Pwd:   "",
				},
			},
		},
	}

	if err := s.initClients(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &s
}

func TestGoRedis(t *testing.T) {
	srv := initSrv()

	kv := srv.ClientContext(0)
	val, err := kv.Client().Get(*kv.Ctx(), "ds.skuids").Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(val)
}
