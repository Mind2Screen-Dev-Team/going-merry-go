package restkey

type RestKey string

func (s RestKey) Name() string {
	switch s {
	case SUCCESS:
		return "SUCCESS"
	case FAILED:
		return "FAILED"
	case INVALID_ARGUMENT:
		return "INVALID_ARGUMENT"
	case NOT_FOUND:
		return "NOT_FOUND"
	case UNAUTHENTICATED:
		return "UNAUTHENTICATED"
	case INTERNAL:
		return "INTERNAL"
	default:
		return "UNKNOWN"
	}
}

// please don't change sequence case
func (s RestKey) Code() int {
	switch s {
	case SUCCESS:
		return 1
	case FAILED:
		return 2
	case INVALID_ARGUMENT:
		return 3
	case NOT_FOUND:
		return 4
	case UNAUTHENTICATED:
		return 5
	case INTERNAL:
		return 6
	default:
		return 0
	}
}

const (
	UNKNOWN          RestKey = "UNKNOWN"
	SUCCESS          RestKey = "SUCCESS"
	FAILED           RestKey = "FAILED"
	INVALID_ARGUMENT RestKey = "INVALID_ARGUMENT"
	NOT_FOUND        RestKey = "NOT_FOUND"
	UNAUTHENTICATED  RestKey = "UNAUTHENTICATED"
	INTERNAL         RestKey = "INTERNAL"

	// add more here...
)
