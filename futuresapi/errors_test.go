package futuresapi

import (
	"errors"
	"testing"

	"github.com/ashexchange/matchclient/v2/types"
)

func TestError(t *testing.T) {
	e := &types.Error{
		Code:    3003,
		Message: "service timeout",
	}

	if !errors.Is(e, ErrServiceTimeout) {
		t.Error("is not a service timeout")
	}
}
