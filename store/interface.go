package store

// IStore is the interface to manage cache data.
type IStore interface {
	// Get
	Get(key interface{}) (interface{}, error)

	// Set .
	Set(key interface{}, value interface{}, options *Options) error

	// Delete .
	Delete(key interface{}) error

	// Invalidate .
	Invalidate(options InvalidateOptions) error

	// Clear remove all cache data
	Clear() error

	// GetType to specify which store your using
	// if your are using pure store, it will be simple "redis" like, otherwise
	// you'll got "wrapper1.wrapper2.redis" to show that how the store wrapped.
	GetType() string
}
