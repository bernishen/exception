package test

import (
	"encoding/json"
	"exception"
	"log"
	"testing"
)

func TestEnum(t *testing.T) {
	log.Print(exception.Information)
	log.Print(exception.Warning)
	log.Print(exception.Error)
	log.Print(exception.FatalError)
}

func TestNewException(t *testing.T) {
	ex := exception.NewException(exception.Error, 1001, "This is an exception.")
	exByte, _ := json.Marshal(ex)
	log.Print(string(exByte))
}

func TestNewExceptionByInfo(t *testing.T) {
	ex := exception.NewExceptionByInfo(exception.Error, exInfoPool[1001])
	exByte, _ := json.Marshal(ex)
	log.Print(string(exByte))
}

var exInfoPool = map[int]*ExInfo{
	1001: {code: 1001, message: "Exception test 1."},
	1002: {code: 1002, message: "Exception test 2."},
}

type ExInfo struct {
	code    int
	message string
}

func (info *ExInfo) Code() int {
	return info.code
}

func (info *ExInfo) Message() string {
	return info.message
}
