package gerror

import (
	"context"
	"fmt"
	"testing"
)

func TestNewErrWithCtx(t *testing.T) {
	ctx := context.WithValue(context.Background(), "traceId", "abcdefghijklmn")
	gerr := NewErrWithCtx(ctx, "error")
	fmt.Println(gerr.TraceID)
}
