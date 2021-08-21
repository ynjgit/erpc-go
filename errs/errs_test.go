package errs_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ynjgit/erpc-go/errs"
)

func TestErr(t *testing.T) {
	var e error
	require.EqualValues(t, errs.CodeOK, errs.Code(e))

	var er *errs.Err
	require.EqualValues(t, errs.CodeOK, errs.Code(er))
	require.EqualValues(t, "OK", errs.Msg(er))

	e = errs.New(111, "haha")
	require.EqualValues(t, 111, errs.Code(e))
	require.EqualValues(t, "haha", errs.Msg(e))
}
