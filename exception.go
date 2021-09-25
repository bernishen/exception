package exception

// Exception : Exception message.
type Exception struct {
	Scope   ExceptionScope `json:"scope"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
}

type IExceptionInfo interface {
	Code() int
	Message() string
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

// NewExceptionByInfo is init a error info entity.
func NewExceptionByInfo(scope ExceptionScope, info IExceptionInfo) *Exception {
	code, msg := readInfo(info)
	ex := NewException(scope, code, msg)
	return ex
}

func readInfo(info IExceptionInfo) (int, string) {
	code, ok1 := func(i IExceptionInfo) (int, bool) {
		defer func() { recover() }()
		return i.Code(), true
	}(info)
	msg, ok2 := func(i IExceptionInfo) (string, bool) {
		defer func() { recover() }()
		return i.Message(), true
	}(info)

	if !ok2 {
		return -1, "未正确获取异常消息｜No exception message was receive"
	}
	if ok2 && !ok1 {
		return -1, "[未正确获取异常代码｜No exception code was receive]" + msg
	}
	return code, msg
}

func (s *ExceptionScope) ToString() string {
	switch *s {
	case Information:
		return "Information"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case FatalError:
		return "FatalError"
	default:
		return ""
	}
}
