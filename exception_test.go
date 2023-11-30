package exception

import (
	"encoding/json"
	"fmt"
	"github.com/bernishen/exception/domain/types/msgscope"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	RegisterLogger(&testLog{})

	a := New(msgscope.Information, -99, "测试创建异常 | Test creation exception")
	a1, err := json.Marshal(a)
	if err != nil {
		t.Log(err.Error())
	}
	builder := strings.Builder{}
	builder.Write(a1)
	t.Log("Entity: " + builder.String())
}

type testLog struct {
	msg string
}

func init() {
	RegisterLogger(&testLog{})
}

func (r *testLog) Log(scope msgscope.MessageScope, code int, msg string) {
	r.msg = msg
	fmt.Printf("[%v<code:%v>] %s \n\r", scope.String(), code, msg)
}
