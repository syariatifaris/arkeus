package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetErrorAndCode(t *testing.T) {
	cases := []struct {
		codes        Codes
		expectString string
		expectInt    int
	}{
		{
			codes:        NoDataFound,
			expectString: "No data found",
			expectInt:    http.StatusBadRequest,
		},
		{
			codes:        DatabaseTypeNotExists,
			expectString: "Database type is not exists",
			expectInt:    http.StatusInternalServerError,
		},
		{
			codes:        RedisNotExists,
			expectString: "Internal server error",
			expectInt:    http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		actualString, actualInt := tc.codes.GetErrorAndCode()
		assert.Equal(t, tc.expectString, actualString)
		assert.Equal(t, tc.expectInt, actualInt)
	}
}

func TestErr(t *testing.T) {
	code := new(Codes)

	err := code.Err()
	if err == nil {
		t.Fatal("No Error Found")
	}
}
