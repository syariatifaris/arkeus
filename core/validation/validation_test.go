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

	"github.com/stretchr/testify/assert"
)

func TestValidationImpl(t *testing.T) {
	v := NewValidator()

	v.Length("tokopedia", 8).SetErrorMessage("Pesan harus sepanjang 9 karakter")
	v.MinSize("PasswordKu", 5).SetErrorMessage("Password harus di atas 5 karakter")
	v.MaxSize("PasswordKu", 7).SetErrorMessage("Password harus di bawah 7 karakter")
	v.Email("syariati.farisxgmail.com").SetErrorMessage("Email harus valid")
	v.Phone("+6281315123964").SetErrorMessage("Phone Number Must be Valid")
	v.URL("http://tokopedia.com/").SetErrorMessage("Url harus valid")
	v.Required("abc").SetErrorMessage("Value is required")
	v.Min(10, 1).SetErrorMessage("Should be more than 1")
	v.Max(9, 10).SetErrorMessage("Should be less than 10")
	v.Equals("abc", "abc").SetErrorMessage("Should be equal")

	_, errs := v.Validate()

	if errs != nil {
		for _, err := range errs {
			assert.NotNil(t, err)
		}
	}
}
