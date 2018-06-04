package log

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syariatifaris/arkeus/core/errors"
)

func TestLog(t *testing.T) {
	testSetGetLevel(t)
	testSetConfig(t)
	testReopenNil(t)
	testgetFileAndLine(t)
	testOthers(t)
}

func testSetGetLevel(t *testing.T) {
	cases := []struct {
		level          string
		expectedResult string
	}{
		{
			level:          "panic",
			expectedResult: "PANIC",
		},
		{
			level:          "fatal",
			expectedResult: "FATAL",
		},
		{
			level:          "error",
			expectedResult: "ERROR",
		},
		{
			level:          "warning",
			expectedResult: "WARNING",
		},
		{
			level:          "info",
			expectedResult: "INFO",
		},
		{
			level:          "debug",
			expectedResult: "DEBUG",
		},
		{
			level:          "default",
			expectedResult: "INFO",
		},
	}
	for _, tc := range cases {
		SetLevel(tc.level)
		assert.Equal(t, tc.expectedResult, GetLevel())
	}
}

func testSetConfig(t *testing.T) {
	config := Config{
		LogLevel:     "",
		ErrorLogPath: "",
	}
	SetConfig(config)
	assert.Equal(t, "INFO", GetLevel())
}

func testReopenNil(t *testing.T) {
	reopen(1, "")
}

func testgetFileAndLine(t *testing.T) {
	file, _ := getFileAndLine()

	filePath := formatFilePath(file)
	assert.Equal(t, filePath, "log_test.go")
}

func testOthers(t *testing.T) {
	param := "test"

	Info(param)
	Infoln(param)
	Infof(param, "log infof")
	Print(param)
	Println(param)
	Printf(param, "log printf")
	Debug(param)
	Debugln(param)
	Debugf(param, "log debugf")
	Warn(param)
	Warnln(param)
	Warnf(param, "log warnf")
	Error(param)
	Errorln(param)
	Errorf(param, "log errorf")
	Errors(fmt.Errorf("Test Error"))
	Errors(errors.New("Some error", errors.New("Some error2")))

	//can't be test, because will cause error
	// Fatal(param)
	// Fatalln(param)
	// Fatalf(param, "log errorf")

	WithFields(Fields{
		"error": fmt.Errorf("Testing error"),
	}).Error("Testing")

}
