package errno

const (
	// SUCCESS - 200: OK.
	SUCCESS errorno = iota + 100001

	// ErrUnknown - 500: Internal server error.
	ErrUnknown
)
