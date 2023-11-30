package exception

import "github.com/bernishen/exception/domain/types/msgscope"

type Exception struct {
	Scope   msgscope.MessageScope `json:"scope"`
	Code    int                   `json:"code"`
	Message string                `json:"message"`
}

type OnCreatedEvent func(ex *Exception)

var OnCreated []OnCreatedEvent

// NewException is init a error info entity.
func NewException(scope msgscope.MessageScope, code int, msg string) *Exception {
	ex := Exception{scope, code, msg}
	if OnCreated != nil || len(OnCreated) == 0 {
		return &ex
	}
	for _, event := range OnCreated {
		if event == nil {
			continue
		}
		event(&ex)
	}
	return &ex
}

// ResetCode : Re write the innercode of this exception.
func (ex *Exception) ResetCode(newCode int) *Exception {
	ex.Code = newCode
	return ex
}
