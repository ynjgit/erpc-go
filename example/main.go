package main

import (
	erpc "github.com/ynjgit/erpc-go"
	_ "github.com/ynjgit/erpc-go/core"
	_ "github.com/ynjgit/erpc-go/interceptor/debuglog"
	_ "github.com/ynjgit/erpc-go/interceptor/recovery"
	"github.com/ynjgit/erpc-go/log"
	_ "github.com/ynjgit/erpc-go/log/console"
	_ "github.com/ynjgit/erpc-go/log/file"

    rpc "test/helloworld"
)

func main() {
    s, err := erpc.NewServer()
    if err != nil {
        log.Fatal(err)
    }

    rpc.RegisterGreeterRPC(s, &greeterImpl{})

    err = s.Start()
    if err != nil {
        log.Fatal(err)
    }
}