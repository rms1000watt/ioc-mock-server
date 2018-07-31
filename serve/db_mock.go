package serve

import log "github.com/sirupsen/logrus"

// Mock is the struct for mock backend
type Mock struct {
	KV map[string]string
}

// Get will retrieve a val from the map storage
func (m Mock) Get(key string) (val string, err error) {
	return m.KV[key], nil
}

// Set will set a value into the map storage
func (m Mock) Set(key, val string) (err error) {
	m.KV[key] = val
	return
}

func newMock(cfg Config) (db DB, err error) {
	log.Debug("Creating Mock DB conn")

	db = Mock{
		KV: map[string]string{},
	}
	return
}
