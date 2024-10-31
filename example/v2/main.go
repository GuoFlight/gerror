package main

import (
	"context"
	"fmt"
	"github.com/GuoFlight/gerror"
	"log"
	"strings"
)

func PrintErrLog(text string) *gerror.Gerr {
	// gerr := NewErr(text)
	// gerr := SetSkip(2).NewErr(text)
	// gerr := SetTraceId("IamTraceId").SetSkip(2).NewErr(text)
	// gerr := SetSkip(2).SetTraceId("IamTraceId").NewErr(text)
	gerr := gerror.SetSkip(2).SetTraceIdByCtx(context.WithValue(context.Background(), gerror.KeyNameTraceId, "IamTraceId")).NewErr(text)
	fmt.Println(strings.Repeat("-----------------------", 3))
	fmt.Println(gerr.ErrFile)
	fmt.Println(gerr.ErrLine)
	fmt.Println(gerr.ErrFunc)
	fmt.Println(gerr.TraceID)
	fmt.Println(gerr.Error())
	fmt.Println(strings.Repeat("-----------------------", 3))
	return gerr
}
func DoSomething() *gerror.Gerr {
	errMsg := "I am error"     // 模拟产生了错误
	return PrintErrLog(errMsg) // 打印错误日志，并返回错误
}

func main() {
	gerr := DoSomething()
	if gerr != nil {
		log.Fatal(gerr.Error())
	}
}
