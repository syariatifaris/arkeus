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

package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetV4UUID(t *testing.T) {
	id, err := GetV4UUID()
	assert.NotNil(t, id)
	assert.NoError(t, err)
}
