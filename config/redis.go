package config

import (
	"os"

	"codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/redis"
)

var (
	Redis = map[string]redis.RedisOptions{
		"cache": {
			//Host: os.Getenv("redis_host"),
			//Port: os.Getenv("redis_port"),
			//Pwd:  os.Getenv("redis_password"),
			Host: os.Getenv("127.0.0.1"),
			Port: os.Getenv("6379"),
			Pwd:  os.Getenv(""),
			Db:   0,
		},
	}
)
