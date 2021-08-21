package recovery_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ynjgit/erpc-go/interceptor"
)

func panicHandle(ctx context.Context, req interface{}, rsp interface{}) error {
	a := 1
	b := 0
	c := a / b
	if c != 0 {
		c = 0
	}
	return nil
}

func TestRecovery(t *testing.T) {
	var ch interceptor.Chain
	i := interceptor.GetServerInterceptor("recovery")
	require.NotNil(t, i)
	ch = append(ch, i)

	ctx := context.Background()
	err := ch.Handle(ctx, nil, nil, panicHandle)
	require.Nil(t, err)
}
