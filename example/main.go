package main

import (
	"errors"
	"fmt"
	"github.com/GuoFlight/gerror"
	"log"
)

func PrintLog(err *gerror.Gerr) *gerror.Gerr {
	log.Println(err.Error())
	return err
}
func DoSomething() *gerror.Gerr {
	err := errors.New("错误") //模拟错误
	if err != nil {
		return gerror.NewErr(err.Error())
	}
	return nil
}
func main() {
	//示例1
	fmt.Println("示例1：")
	err1 := gerror.NewErr("错误1")
	if err1 != nil {
		fmt.Println(err1.TraceID)
		fmt.Println(err1.ErrFile)
		fmt.Println(err1.ErrFunc)
		fmt.Println(err1.ErrLine)
		fmt.Println(err1.Error())
	}
	fmt.Println()

	//示例2
	fmt.Println("示例2：")
	err2 := gerror.NewErr("错误2", err1.TraceID)
	fmt.Println(err2.TraceID)
	fmt.Println()

	//示例3
	fmt.Println("示例3：")
	err3 := PrintLog(gerror.NewErr("错误3", err2.TraceID))
	fmt.Println(err3)
	fmt.Println()

	//示例4
	fmt.Println("示例4：")
	err4 := DoSomething()
	fmt.Println(err4)
}
