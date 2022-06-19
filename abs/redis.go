package abs

import (
	"go_test/config"

	rds "codeup.aliyun.com/5f69c1766207a1a8b17fda8e/sanhe_library/redis"
	"github.com/go-redis/redis"
)

type Redis struct {
	rds.Redis
}

var cache rds.RdsCollector

func (r *Redis) GetCache() *redis.Client {
	return r.GetClient(&cache, config.Redis["cache"])
}
