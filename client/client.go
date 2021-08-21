package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ynjgit/erpc-go/config"
	"github.com/ynjgit/erpc-go/errs"
	"github.com/ynjgit/erpc-go/interceptor"
	"github.com/ynjgit/erpc-go/protocol"
)

// Client the client inerface, Call function to do remoute call
type Client interface {
	Call(ctx context.Context, req interface{}, rsp interface{}, opts ...Option) error
}

// DefaultClient the default client
var DefaultClient = New()

// New new http client
func New() Client {
	return &httpClient{
		client: &http.Client{},
	}
}

type httpClient struct {
	client *http.Client
}

// Call http remoute call
func (c *httpClient) Call(ctx context.Context, req interface{}, rsp interface{}, opts ...Option) error {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}

	cliConfig := config.GetClientConfig()
	for _, name := range cliConfig.Interceptor {
		i := interceptor.GetClientInterceptor(name)
		options.Intercetors = append(options.Intercetors, i)
	}

	var remoteConfig *config.RemoteConfig
	for _, remote := range cliConfig.Remote {
		if remote.Name == options.SVCName {
			remoteConfig = remote
		}
	}
	if remoteConfig != nil {
		options.Target = remoteConfig.Target
		options.TimeOut = time.Duration(remoteConfig.Timeout) * time.Millisecond
	}

	erpcCtx := protocol.GetCtx(ctx)
	erpcCtx.ClientRPCName = options.RPCName

	handler := func(ctx context.Context, req interface{}, rsp interface{}) error {

		reqBody, err := protocol.Marshal(req)
		if err != nil {
			return errs.New(errs.CodeClientMarshalFail, err.Error())
		}

		if options.TimeOut > 0 {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, options.TimeOut)
			defer cancel()
		}

		url := fmt.Sprintf("%s%s", options.Target, options.RPCName)
		request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
		if err != nil {
			return errs.New(errs.CodeClientCallFail, err.Error())
		}

		response, err := c.client.Do(request)
		if err != nil {
			return errs.New(errs.CodeClientCallFail, err.Error())
		}
		defer response.Body.Close()

		if response.StatusCode >= http.StatusMultipleChoices {
			return errs.New(response.StatusCode, response.Status)
		}

		rspBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return errs.New(errs.CodeClientCallFail, err.Error())
		}

		err = protocol.Unmarshal(rspBody, rsp)
		if err != nil {
			return errs.New(errs.CodeClientUnmarshalFail, err.Error())

		}
		return nil
	}

	return options.Intercetors.Handle(ctx, req, rsp, handler)
}
