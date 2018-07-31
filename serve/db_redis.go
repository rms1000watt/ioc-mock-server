package serve

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// Redis is the struct to connect to redis
type Redis struct {
	conn *redis.Client
}

// Get will get a key from redis
func (r Redis) Get(key string) (val string, err error) {
	// TODO: Check for healthy conn and reconnected if err

	return r.conn.Get(key).Result()
}

// Set will set a key value pair into redis
func (r Redis) Set(key, val string) (err error) {
	// TODO: Check for healthy conn and reconnected if err

	return r.conn.Set(key, val, 0).Err()
}

func newRedis(cfg Config) (db DB, err error) {
	log.Debug("Creating Redis DB conn")

	r := Redis{}
	r.conn = redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	if _, err = r.conn.Ping().Result(); err != nil {
		log.Debug("Failed pinging redis: ", err)
		return
	}

	db = r
	return
}
