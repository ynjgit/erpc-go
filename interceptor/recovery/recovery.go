package recovery

import (
	"context"
	"runtime"

	"github.com/ynjgit/erpc-go/interceptor"
	"github.com/ynjgit/erpc-go/log"
)

func init() {
	interceptor.Register("recovery", recovery, nil)
}

// recovery interceptor
func recovery(ctx context.Context, req interface{}, rsp interface{}, f interceptor.Handler) error {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 1024)
			n := runtime.Stack(buf, false)
			buf = buf[:n]

			log.ErrorContextf(ctx, "panic:%v\ntrace:%s\n", r, buf)
		}
	}()

	return f(ctx, req, rsp)
}
