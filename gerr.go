package gerror

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"runtime"
)

var KeyNameTraceId = "traceId"

type Gerr struct {
	s       string // 保存错误信息
	TraceID string
	ErrFile string
	ErrFunc string
	ErrLine int
}

type Entry struct {
	Skip    int
	TraceID string
}

// Error 兼容原生error
func (g Gerr) Error() string {
	return g.s
}

// SetSkip 设置caller的skip值
func SetSkip(skip int) *Entry {
	var entry Entry
	return entry.SetSkip(skip)
}
func (e *Entry) SetSkip(skip int) *Entry {
	e.Skip = skip
	return e
}

// SetTraceId 设置TraceId
func SetTraceId(traceId string) *Entry {
	var entry Entry
	return entry.SetTraceId(traceId)
}
func (e *Entry) SetTraceId(traceId string) *Entry {
	e.TraceID = traceId
	return e
}

// SetTraceIdByCtx 从context中获取并设置traceId
func SetTraceIdByCtx(ctx context.Context) *Entry {
	var entry Entry
	return entry.SetTraceIdByCtx(ctx)
}
func (e *Entry) SetTraceIdByCtx(ctx context.Context) *Entry {
	traceId, ok := ctx.Value(KeyNameTraceId).(string)
	if !ok {
		traceId = uuid.NewV4().String()
	}
	e.TraceID = traceId
	return e
}

// NewErr 新建Err
func (e *Entry) NewErr(text string) *Gerr {
	// 获取调用者信息
	skip := e.Skip
	if skip == 0 {
		skip = 1
	}
	errFuncName := ""
	pc, errFile, errLine, ok := runtime.Caller(skip)
	if ok {
		errFuncName = runtime.FuncForPC(pc).Name()
	}

	// 获取traceId
	traceId := e.TraceID
	if traceId == "" {
		traceId = uuid.NewV4().String()
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
func NewErr(text string) *Gerr {
	traceID := uuid.NewV4().String()
	var entry Entry
	return entry.SetSkip(2).SetTraceId(traceID).NewErr(text)
}
