package database

import (
	"fmt"
	"sync"

	"github.com/fwidjaya20/symphonic/facades"
	"github.com/redis/go-redis/v9"
)

var (
	instanceRedis *redis.Client
	syncRedis     sync.Once
)

func Redis() *redis.Client {
	syncRedis.Do(func() {
		instanceRedis = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf(
				"%s:%d",
				facades.Config().GetString("database.connections.redis.host"),
				facades.Config().GetInt("database.connections.redis.port"),
			),
			Password: facades.Config().GetString("database.connections.redis.password"),
			DB:       facades.Config().GetInt("database.connections.redis.database"),
		})
	})

	return instanceRedis
}
