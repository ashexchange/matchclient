package types

import (
	"testing"
)

func TestError_As(t *testing.T) {
	var rpcErr error = &Error{
		Code:    1000,
		Message: "error",
	}

	e := FromError(rpcErr)

	t.Logf("%s", e)
}
