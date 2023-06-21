package cache

import "time"

type Cache interface {
	Set(key, value interface{}, ttl time.Duration) bool
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{}) bool
}
