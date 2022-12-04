package store

// Factory defines the storage interface.
type Factory interface {
	Close() error
	
	Users() User
}
