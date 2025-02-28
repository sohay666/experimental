package utils

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

/*
--------------------------------------------------------------------------
| DB   | DESCRIPTION                                                     |
--------------------------------------------------------------------------
|  0   | Handle for cache while user input some keyword search
*/

type DBRedis struct {
	RedisClient *redis.Client
}

var (
	CACHE_VIDEO        = 0
	onceDbRedis        sync.Once
	instanceDBRedis    *DBRedis
	instanceDBRedisMap = map[int]*DBRedis{
		0: nil,
	}
	onceDbRedisMap = map[int]*sync.Once{
		0: nil,
	}
)

type RedisConfig struct {
	Host     string
	DB       int
	Port     int
	Password string
}

// This connection for redis
func GetInstanceRedis(Db int) *redis.Client {
	if onceDbRedisMap[Db] == nil {
		onceDbRedisMap[Db] = &sync.Once{}
	}
	onceDbRedisMap[Db].Do(func() {
		redisInfo := RedisConfig{
			Host:     Config.Service.Redis.Host,
			DB:       Db,
			Port:     Config.Service.Redis.Port,
			Password: Config.Service.Redis.Password,
		}
		logs := fmt.Sprintf("[INFO] Connected to REDIS Host = %s | Db = %d", redisInfo.Host, redisInfo.DB)

		clientConnection := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", redisInfo.Host, redisInfo.Port),
			Password: redisInfo.Password,
			DB:       redisInfo.DB,
		})

		if _, err := clientConnection.Ping(context.Background()).Result(); err != nil {
			logs = "[ERROR] Failed to connect to Redis. Config=" + redisInfo.Host
			log.Println(logs)
			return
		}
		instanceDBRedisMap[Db] = &DBRedis{RedisClient: clientConnection}
	})
	return instanceDBRedisMap[Db].RedisClient
}
