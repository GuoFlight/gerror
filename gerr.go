package gerror

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"runtime"
)

func NewErr(text string, traceID ...string) *Gerr {
	// 获取TraceID
	traceIDTemp := ""
	if len(traceID) > 0 {
		traceIDTemp = traceID[0]
	} else {
		traceIDTemp = uuid.NewV4().String()
	}

	// 获取调用者信息
	errFuncName := ""
	pc, errFile, errLine, ok := runtime.Caller(1)
	if ok {
		errFuncName = runtime.FuncForPC(pc).Name()
	}

	// return
	return &Gerr{
		s:       text,
		TraceID: traceIDTemp,
		ErrFile: errFile,
		ErrFunc: errFuncName,
		ErrLine: errLine,
	}
}

// NewErrWithCtx 从ctx中获取traceId，创建Gerr对象
func NewErrWithCtx(ctx context.Context, text string) *Gerr {
	traceId, ok := ctx.Value("traceId").(string)
	if !ok {
		traceId = uuid.NewV4().String()
	}

	// 获取调用者信息
	errFuncName := ""
	pc, errFile, errLine, ok := runtime.Caller(1)
	if ok {
		errFuncName = runtime.FuncForPC(pc).Name()
	}

	// return
	return &Gerr{
		s:       text,
		TraceID: traceId,
		ErrFile: errFile,
		ErrFunc: errFuncName,
		ErrLine: errLine,
	}
}

type Gerr struct {
	s       string // 保存错误信息
	TraceID string
	ErrFile string
	ErrFunc string
	ErrLine int
}

// 兼容原生error
func (g Gerr) Error() string {
	return g.s
}
