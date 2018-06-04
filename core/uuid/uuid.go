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
	"os/exec"
	"strings"
)

//GetV4UUID will retrieve new universal identifier for computer system
func GetV4UUID() (string, error) {
	//use satori uuid library
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}
