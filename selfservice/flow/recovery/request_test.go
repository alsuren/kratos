package recovery_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/urlx"

	"github.com/ory/kratos/selfservice/flow/recovery"
)

func TestRequest(t *testing.T) {
	must := func(r *recovery.Request, err error) *recovery.Request {
		require.NoError(t, err)
		return r
	}

	u := &http.Request{URL: urlx.ParseOrPanic("http://foo/bar/baz"), Host: "foo"}
	for k, tc := range []struct {
		r         *recovery.Request
		expectErr bool
	}{
		{r: must(recovery.NewRequest(time.Hour, "", u, nil))},
		{r: must(recovery.NewRequest(-time.Hour, "", u, nil)), expectErr: true},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			err := tc.r.Valid()
			if tc.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}

	assert.EqualValues(t, recovery.StateChooseMethod,
		must(recovery.NewRequest(time.Hour, "", u, nil)).State)
}
