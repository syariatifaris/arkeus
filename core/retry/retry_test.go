package retry

import (
	"errors"
	"fmt"
	"testing"

	"net/http"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/syariatifaris/arkeus/core/apicalls"
)

const MaxAttempt = 10

func TestPolicyImpl_Execute(t *testing.T) {
	apicalls.Init(3, false)
	tests := []struct {
		URL                string
		ExpectedResultCode int
		ExpectedError      error
	}{
		{
			URL:                "http://localhost/hook/incoming/fist",
			ExpectedResultCode: http.StatusOK,
			ExpectedError:      nil,
		},
		{
			URL:                "http://localhost/hook/incoming/second",
			ExpectedResultCode: http.StatusInternalServerError,
			ExpectedError:      nil,
		},
		//negative test case
		{
			URL:           "http://localhost/hook/incoming/third",
			ExpectedError: errors.New("gock: cannot match any request"),
		},
	}

	for _, tc := range tests {
		gock.New(tc.URL).Reply(tc.ExpectedResultCode).SetError(tc.ExpectedError)
		retryPolicy := NewRetryPolicy(MaxAttempt)
		var actualStatusCode int

		err := retryPolicy.Execute(func() (retry bool, err error) {
			var next = true
			httpAPI := apicalls.HttpAPI{
				URL:    tc.URL,
				Method: http.MethodPost,
			}

			response, err := apicalls.DoRequest(httpAPI)
			if response != nil {
				actualStatusCode = response.StatusCode
				if response.StatusCode == http.StatusOK {
					next = false
				}
			}
			return next, err
		})

		if err == nil {
			assert.Equal(t, tc.ExpectedError, err)
		} else {
			modifiedErrString := fmt.Sprintf("Retrying %d times. Error cause: Post %s: %s", MaxAttempt, tc.URL, tc.ExpectedError.Error())
			assert.Equal(t, modifiedErrString, err.Error())
		}

		assert.Equal(t, tc.ExpectedResultCode, actualStatusCode)
	}
}
