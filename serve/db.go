package serve

import (
	"errors"
	"strings"
)

// DB interface allows redis or mock backend
type DB interface {
	Get(key string) (value string, err error)
	Set(key, value string) (err error)
}

// newDB is the inversion of control (dependency injection)
func newDB(cfg Config) (db DB, err error) {
	switch strings.TrimSpace(strings.ToLower(cfg.DBType)) {
	case "redis":
		return newRedis(cfg)
	case "mock":
		return newMock(cfg)
	}

	return nil, errors.New("Bad dbType")
}
