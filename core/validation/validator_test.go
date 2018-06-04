// Copyright (c) 2018.
// PT.Tokopedia
//
// NOTICE:  All information contained herein is, and remains
// the property of PT.Tokopedia Incorporated and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to PT.Tokopedia Incorporated
// and its suppliers and may be covered, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from PT.Tokopedia Incorporated.

package validation

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestMaxIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		maxValue    int
		isFulfilled bool
	}{
		{
			name:        "positive test case",
			value:       10,
			maxValue:    100,
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       10,
			maxValue:    9,
			isFulfilled: false,
		},
		{
			name:        "negative test case 2",
			value:       "asdf",
			maxValue:    9,
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Max{
				Max: test.maxValue,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestMinIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		minValue    int
		isFulfilled bool
	}{
		{
			name:        "positive test case",
			value:       10,
			minValue:    1,
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       10,
			minValue:    11,
			isFulfilled: false,
		},
		{
			name:        "negative test case 2",
			value:       "asdf",
			minValue:    10,
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Min{
				Min: test.minValue,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestMin64IsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		minValue    int64
		isFulfilled bool
	}{
		{
			name:        "positive test case",
			value:       1000000000000001,
			minValue:    1000000000000000,
			isFulfilled: false,
		},
		{
			name:        "negative test case 1",
			value:       1000000000000001,
			minValue:    1000000000000002,
			isFulfilled: false,
		},
		{
			name:        "negative test case 2",
			value:       "asdf",
			minValue:    1000000000000001,
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Min64{
				Min: test.minValue,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestRangeIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		minBound    int
		maxBound    int
		isFulfilled bool
	}{
		{
			name:        "positive test case",
			value:       10,
			minBound:    1,
			maxBound:    11,
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       10,
			minBound:    11,
			maxBound:    13,
			isFulfilled: false,
		},
		{
			name:        "negative test case 2",
			value:       "asdf",
			minBound:    11,
			maxBound:    13,
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Range{
				Min: Min{
					Min: test.minBound,
				},
				Max: Max{
					Max: test.maxBound,
				},
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestRequiredIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       time.Now(),
			isFulfilled: true,
		},
		{
			name:        "positive test case 2",
			value:       10,
			isFulfilled: true,
		},
		{
			name:        "positive test case 3",
			value:       "asdf",
			isFulfilled: true,
		},
		{
			name: "positive test case 4",
			value: struct {
			}{},
			isFulfilled: true,
		},
		{
			name:        "positive test case 5",
			value:       true,
			isFulfilled: true,
		},
		{
			name:        "positive test case int64",
			value:       int64(999999),
			isFulfilled: true,
		},
		{
			name:        "positive test case int8",
			value:       int8(9),
			isFulfilled: true,
		},
		{
			name:        "positive test case int16",
			value:       int16(9),
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       nil,
			isFulfilled: false,
		},
		{
			name:        "negative test case 2",
			value:       0,
			isFulfilled: false,
		},
		{
			name:        "test zero int64",
			value:       int64(0),
			isFulfilled: false,
		},
		{
			name:        "test zero int",
			value:       int(0),
			isFulfilled: false,
		},
		{
			name:        "test zero int8",
			value:       int8(0),
			isFulfilled: false,
		},
		{
			name:        "test zero int16",
			value:       int16(0),
			isFulfilled: false,
		},
		{
			name:        "test zero int32",
			value:       int32(0),
			isFulfilled: false,
		},
		{
			name:        "test empty string",
			value:       "",
			isFulfilled: false,
		},
		{
			name:        "test empty time",
			value:       time.Time{},
			isFulfilled: false,
		},
		{
			name:        "test empty slice",
			value:       []int{},
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Required{}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestMinSizeIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		minSize     int
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "asdf",
			minSize:     1,
			isFulfilled: true,
		},
		{
			name: "positive test case 2",
			value: []string{
				"a", "b",
			},
			minSize:     1,
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       "asdf",
			minSize:     11,
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := MinSize{
				Min: test.minSize,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestMaxSizeIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		maxSize     int
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "asdf",
			maxSize:     5,
			isFulfilled: true,
		},
		{
			name: "positive test case 2",
			value: []string{
				"a", "b",
			},
			maxSize:     3,
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       "asdf",
			maxSize:     1,
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := MaxSize{
				Max: test.maxSize,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestLengthIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		length      int
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "asdf",
			length:      4,
			isFulfilled: true,
		},
		{
			name: "positive test case 2",
			value: []string{
				"a", "b",
			},
			length:      2,
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       "asdf",
			length:      1,
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Length{
				N: test.length,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestMatchIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		compare     string
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "syariati.faris@gmail.com",
			compare:     "syariati.faris@gmail.com",
			isFulfilled: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Match{
				Regexp: emailPattern,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestEmailIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "syariati.faris@gmail.com",
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       "syariati.farisxgmail.com",
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Email{
				Match: Match{
					Regexp: emailPattern,
				},
			}
			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestDomainIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "syariati.com",
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       "syariatidotcom",
			isFulfilled: false,
		},
		{
			name: "negative test case 3",
			value: "syariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidot" +
				"comsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidot" +
				"comsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidot" +
				"comsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcomsyariatidotcom",
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Domain{
				Regexp: domainPattern,
			}

			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestURLIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "http://tokopedia.com",
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       "tokopedia(dot)com",
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := URL{
				Domain: Domain{
					Regexp: domainPattern,
				},
			}
			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}

func TestEqualsIsFulfilled(t *testing.T) {
	tests := []struct {
		name        string
		value       interface{}
		compare     interface{}
		isFulfilled bool
	}{
		{
			name:        "positive test case 1",
			value:       "abc",
			compare:     "abc",
			isFulfilled: true,
		},
		{
			name:        "negative test case 1",
			value:       "cdf",
			compare:     "abc",
			isFulfilled: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := Equals{
				CompareObj: test.compare,
			}
			ok := m.IsFulfilled(test.value)
			msg := m.GetDefaultMessage()
			assert.Equal(t, test.isFulfilled, ok)
			assert.NotNil(t, msg)
		})
	}
}
