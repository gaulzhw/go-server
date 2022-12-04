package store

var client Factory

// Factory defines the storage interface.
type Factory interface {
	Users() User

	Close() error
}

// Client return the store client instance.
func Client() Factory {
	return client
}

// SetClient set the iam store client.
func SetClient(factory Factory) {
	client = factory
}
