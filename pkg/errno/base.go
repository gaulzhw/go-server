package errno

const (
	// Success - 200: OK.
	Success errorno = iota + 100001

	// ErrUnknown - 500: Internal server error.
	ErrUnknown

	// ErrNotFound - 404: Not found error.
	ErrNotFound
)
