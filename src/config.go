package src

import "github.com/springmove/sptty"

type RedisEntry struct {
	Entry string `yaml:"Entry"`
	Pwd   string `yaml:"Pwd"`

	// only for redis
	DB int `yaml:"DB"`
}

type Config struct {
	sptty.BaseConfig
	Configs []RedisEntry `yaml:"Configs"`
}

func (s *Config) ConfigName() string {
	return ServiceGoRedis
}
