package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShopAPI(t *testing.T) {
	testErrs(t)
	testNew(t)
	//testMessages(t)
	testMatch(t)
	testErrorAndHTTPCode(t)
}

func testErrs(t *testing.T) {
	cases := []struct {
		err    *Errs
		expect *Errs
	}{
		{
			err: New("New error without anything"),
			expect: &Errs{
				err: errors.New("New error without anything"),
			},
		},
		{
			err: New("Error with fields", Fields{"first": "two", "satu": 2}),
			expect: &Errs{
				err:    errors.New("Error with fields"),
				fields: Fields{"first": "two", "satu": 2},
			},
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.err, tc.expect)
	}
}

func testNew(t *testing.T) {
	cases := []struct {
		actual       *Errs
		expectResult interface{}
		typeError    string
	}{
		{
			actual:       New("Some error", errors.New("Some error2")),
			expectResult: errors.New("Some error2"),
		},
		{
			actual:       New(New("Some error"), InternalServerError),
			expectResult: errors.New("Internal server error"),
		},
		{
			actual:       New(Fields{"error1": "error11", "error2": "error22"}),
			expectResult: Fields{"error1": "error11", "error2": "error22"},
			typeError:    "fields",
		},
		{
			actual:       New([]string{"error1", "error2"}),
			expectResult: []string{"error1", "error2"},
			typeError:    "messages",
		},
		{
			actual:       New(1),
			expectResult: errors.New("Unknown error"),
		},
		{
			actual: &Errs{
				traces: []string{"1", "2"},
				file:   "abc",
				line:   1,
			},
			typeError: "others",
		},
	}
	SetRuntimeOutput(true)
	assert.Equal(t, true, IsRuntimeEnabled())
	for _, tc := range cases {
		switch tc.typeError {
		case "":
			assert.Equal(t, tc.expectResult.(error).Error(), tc.actual.Error())
		case "fields":
			assert.Equal(t, tc.expectResult.(Fields), tc.actual.GetFields())
		case "messages":
			assert.Equal(t, tc.expectResult.([]string), tc.actual.GetMessages())
		case "others":
			msg := "message1"
			tc.actual.SetMessage(msg)
			assert.Equal(t, msg, tc.actual.GetMessage())
			assert.Equal(t, tc.actual.GetTrace(), tc.actual.traces)
			expectFile, expectLine := tc.actual.GetFileAndLine()
			assert.Equal(t, expectFile, tc.actual.file)
			assert.Equal(t, expectLine, tc.actual.line)
		}
	}
}

func testMatch(t *testing.T) {
	cases := []struct {
		err1        error
		err2        error
		expectMatch bool
	}{
		{
			err1:        New(errors.New("This is new error")),
			err2:        nil,
			expectMatch: false,
		},
		{
			err1:        New(errors.New("This is new error")),
			err2:        errors.New("This is new error"),
			expectMatch: true,
		},
		{
			err1:        nil,
			err2:        errors.New("Something is different"),
			expectMatch: false,
		},
		{
			err1:        New(errors.New("This is new error")),
			err2:        New(errors.New("This is new error")),
			expectMatch: true,
		},
		{
			err1:        nil,
			err2:        nil,
			expectMatch: true,
		},
	}

	for _, tc := range cases {
		actual := Match(tc.err1, tc.err2)
		assert.Equal(t, tc.expectMatch, actual)
	}
}

func testErrorAndHTTPCode(t *testing.T) {
	tests := []struct {
		err      error
		expected int
	}{
		{
			err:      errors.New("Testing Error"),
			expected: http.StatusInternalServerError,
		},
		{
			err:      New(errors.New("This is new error")),
			expected: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		_, errCode := ErrorAndHTTPCode(tc.err)
		if errCode != tc.expected {
			t.Fatal("Wrong Error Code")
		}
	}
}
