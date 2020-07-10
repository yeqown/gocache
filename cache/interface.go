package cache

import (
	"github.com/eko/gocache/codec"
	"github.com/eko/gocache/store"
)

// ICache represents the interface for all caches (aggregates, metric, memory, redis, ...)
type ICache interface {
	Get(key interface{}) (interface{}, error)
	Set(key, object interface{}, options *store.Options) error
	Delete(key interface{}) error
	Invalidate(options store.InvalidateOptions) error
	Clear() error
	GetType() string

	// GetOrSet(key interface{}, value interface{}) (interface{}, error)
	// GetStore() store.IStore
}

// SetterCacheInterface represents the interface for caches that allows
// storage (for instance: memory, redis, ...)
type SetterCacheInterface interface {
	ICache

	GetCodec() codec.CodecInterface
}
