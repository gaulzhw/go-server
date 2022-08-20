package store

//go:generate mockgen -self_package=github.com/gaulzhw/go-server/internal/server/store -destination mock_store.go -package store github.com/gaulzhw/go-server/internal/server/store Factory,UserStore,SecretStore,PolicyStore

var client Factory

// Factory defines the server platform storage interface.
type Factory interface {
	Users() UserStore
	Posts() PostStore
	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the server store client.
func SetClient(factory Factory) {
	client = factory
}
