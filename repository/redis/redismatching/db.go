package redismatching

import "github.com/mobin-alz/gameapp/adapter/redis"

type DB struct {
	adapter redis.Adapter
}

func New(adapter redis.Adapter) DB {
	return DB{
		adapter: adapter,
	}
}
