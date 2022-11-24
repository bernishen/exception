package messageScope

type MessageScope int

const (
	// Information : Is not a exception.
	Information MessageScope = 1 << iota
	// Warning : Is a exception, but program can continue to run.
	Warning
	// Error : This will cause get result is not ture.
	Error
	// FatalError ::This will cause the run to stop.
	FatalError
)

func (scope *MessageScope) String() string {
	switch *scope {
	case Information:
		return "Information"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case FatalError:
		return "FatalError"
	default:
		return "undefined"
	}
}
