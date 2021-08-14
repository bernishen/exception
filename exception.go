package exception

// Exception : Exception message.
type Exception struct {
	Scope ExceptionScope
	Code  int
	Message string
}

// ExceptionScope : A exception's scope.
// this element's :
// Information is 0;
// Warning is 1;
// Error is 2.
type ExceptionScope byte

type OnCreatedEvent func(ex *Exception)

var OnCreated OnCreatedEvent

const (
	// Information : Is not a exception.
	Information ExceptionScope = 1 << iota
	// Warning : Is a exception, but program can continue to run.
	Warning
	// Error : This will cause get result is not ture.
	Error
	// FatalError ::This will cause the run to stop.
	FatalError
)

// ResetCode : Re write the innercode of this exception.
func (ex *Exception) ResetCode(newCode int) *Exception {
	ex.Code = newCode
	return ex
}

// NewException is init a error info entity.
func NewException(scope ExceptionScope, code int, msg string) *Exception {
	ex := Exception{scope, code, msg}
	if OnCreated != nil {
		OnCreated(&ex)
	}
	return &ex
}
