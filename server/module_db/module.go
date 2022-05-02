package module_db

import (
	"github.com/go-redis/redis/v8"
	"github.com/topfreegames/pitaya/modules"
	"server/module_db/model"
)

type RedisModule struct {
	modules.Base
	rdb *redis.Client
}

func NewRedisModule() *RedisModule {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisModule{
		rdb: rdb,
	}
}

func (this *RedisModule) QueryGuestUser(appId string, imei string) (*model.User, error) {
	return nil, nil
}
