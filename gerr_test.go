package gerror

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func MyTest1(text string) {
	fmt.Println(strings.Repeat("-----------------------", 3))
	// gerr := NewErr(text)
	// gerr := SetSkip(2).NewErr(text)
	// gerr := SetTraceId("IamTraceId").SetSkip(2).NewErr(text)
	// gerr := SetSkip(2).SetTraceId("IamTraceId").NewErr(text)
	gerr := SetSkip(2).SetTraceIdByCtx(context.WithValue(context.Background(), KeyNameTraceId, "IamTraceId")).NewErr(text)
	fmt.Println(gerr.ErrFile)
	fmt.Println(gerr.ErrLine)
	fmt.Println(gerr.ErrFunc)
	fmt.Println(gerr.TraceID)
	fmt.Println(gerr.Error())
	fmt.Println(strings.Repeat("-----------------------", 3))
}
func TestGerr(t *testing.T) {
	MyTest1("test")
}
