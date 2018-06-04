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
