package extension

import (
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/vmihailenco/msgpack"
)

var _ cache.ICache = marshalStoreWrapper{}

type marshalStoreWrapper struct {
	store store.IStore
}

// NewMarshalStoreWrapper
func NewMarshalStoreWrapper(oriStore store.IStore) *marshalStoreWrapper {
	return &marshalStoreWrapper{
		store: oriStore,
	}
}

// Get obtains a value from cache and unmarshal value with given object
func (w marshalStoreWrapper) Get(key interface{}, returnObj interface{}) (interface{}, error) {
	result, err := w.store.Get(key)
	if err != nil {
		return nil, err
	}

	switch result.(type) {
	case []byte:
		err = msgpack.Unmarshal(result.([]byte), returnObj)

	case string:
		err = msgpack.Unmarshal([]byte(result.(string)), returnObj)
	}

	if err != nil {
		return nil, err
	}

	return returnObj, nil
}

// Set sets a value in cache by marshaling value
func (w marshalStoreWrapper) Set(key, object interface{}, options *store.Options) error {
	bytes, err := msgpack.Marshal(object)
	if err != nil {
		return err
	}

	return w.store.Set(key, bytes, options)
}

// Delete removes a value from the cache
func (w marshalStoreWrapper) Delete(key interface{}) error {
	return w.store.Delete(key)
}

// Invalidate invalidate cache values using given options
func (w marshalStoreWrapper) Invalidate(options store.InvalidateOptions) error {
	return w.store.Invalidate(options)
}

// Clear reset all cache data
func (w marshalStoreWrapper) Clear() error {
	return w.store.Clear()
}

func (w marshalStoreWrapper) GetType() string {
	panic("implement me")
}
