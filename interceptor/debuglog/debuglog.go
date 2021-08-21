package debuglog

import (
	"context"
	"time"

	"github.com/ynjgit/erpc-go/interceptor"
	"github.com/ynjgit/erpc-go/log"
	"github.com/ynjgit/erpc-go/protocol"
)

func init() {
	interceptor.Register("debuglog", debuglogServer, debuglogClient)
}

// the server debuglog interceptor
func debuglogServer(ctx context.Context, req interface{}, rsp interface{}, f interceptor.Handler) error {
	erpcCtx := protocol.GetCtx(ctx)
	t := time.Now()
	err := f(ctx, req, rsp)
	cost := time.Since(t)
	log.DebugContextf(ctx, "server handle ServerRPCName:%s, req:%+v, rsp:%+v, cost:%v err:%v", erpcCtx.ServerRPCName, req, rsp, cost, err)
	return err
}

// the client debuglog interceptor
func debuglogClient(ctx context.Context, req interface{}, rsp interface{}, f interceptor.Handler) error {
	erpcCtx := protocol.GetCtx(ctx)
	t := time.Now()
	err := f(ctx, req, rsp)
	cost := time.Since(t)
	log.DebugContextf(ctx, "client call ClientRPCName:%s, req:%+v, rsp:%+v, cost:%v err:%v", erpcCtx.ClientRPCName, req, rsp, cost, err)
	return err
}
