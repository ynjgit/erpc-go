package core

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/ynjgit/erpc-go/errs"

	"github.com/labstack/echo/v4"
)

func init() {
	Register("echo", &echoCore{})
}

type echoCore struct {
	echo   *echo.Echo
	option *ServerOption
}

func (ec *echoCore) Serve(ctx context.Context, opts ...ServerOpt) error {
	ec.option = &ServerOption{}
	for _, o := range opts {
		o(ec.option)
	}

	handle := func(c echo.Context) error {
		r := c.Request()
		ctx := r.Context()
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		rspBody, err := ec.option.Handler.Handle(ctx, r.URL.Path, reqBody)
		if err != nil {

			return c.JSON(http.StatusOK, ErrRsp{ErrCode: errs.Code(err), ErrMsg: errs.Msg(err)})
		}

		_, err = c.Response().Write(rspBody)
		return err
	}

	ec.echo = echo.New()
	ec.echo.POST("*", handle)
	return ec.echo.Start(ec.option.ListenAddress)
}

func (ec *echoCore) Shutdown(ctx context.Context) error {
	return ec.echo.Shutdown(ctx)
}
