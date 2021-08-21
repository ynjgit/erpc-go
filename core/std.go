package core

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/ynjgit/erpc-go/errs"
)

func init() {
	Register("std", &stdCore{})
}

type stdCore struct {
	option *ServerOption
	svr    *http.Server
}

func (std *stdCore) Serve(ctx context.Context, opts ...ServerOpt) error {
	std.option = &ServerOption{}
	for _, o := range opts {
		o(std.option)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			code := http.StatusMethodNotAllowed
			http.Error(w, http.StatusText(code), code)
			return
		}

		ctx := r.Context()
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		rspBody, err := std.option.Handler.Handle(ctx, r.URL.Path, reqBody)
		if err != nil {
			http.Error(w, errs.Msg(err), errs.Code(err))
			return
		}

		w.Write(rspBody)
	}
	std.svr = &http.Server{
		Addr:    std.option.ListenAddress,
		Handler: http.HandlerFunc(handler),
	}

	return std.svr.ListenAndServe()
}

func (std *stdCore) Shutdown(ctx context.Context) error {
	return std.svr.Shutdown(ctx)
}
