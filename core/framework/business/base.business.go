package business

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/syariatifaris/arkeus/core/retry"
	"github.com/syariatifaris/arkeus/core/validation"
)

var (
	retryAttempt = 10

	RetryDelay = time.Second
	RedisMax   = time.Second * 2
)

//BaseBusinessModel acts as a business model contract
type BaseBusinessModel interface {
	//GetModel return the business model
	GetModel() (interface{}, error)
	//validate the object
	Validate(obj interface{}) error
}

//BaseBusinessModelWithContext acts as a business model contract by passing a context to operation
type BaseBusinessModelWithContext interface {
	//GetModel return the business model
	GetModelCtx(ctx context.Context) (interface{}, error)
	//validate the object
	Validate(obj interface{}) error
}

//BusinessModel structure
type BusinessModel struct {
	retryAttempt int
	//validation
	validation validation.Validation
}

//validator return the validation.Validation
func (b *BusinessModel) Validator() validation.Validation {
	if b.validation == nil {
		b.validation = validation.NewValidator()
	}

	return b.validation
}

//errorsToString casts array of errors to string
func (*BusinessModel) ErrorsToString(errs []*validation.Error) string {
	var errorString bytes.Buffer

	for i, err := range errs {
		if i == 0 {
			errorString.WriteString(err.Message)
		} else {
			errorString.WriteString(fmt.Sprintf(", %s", err.Message))
		}
	}

	return errorString.String()
}

func (*BusinessModel) NewDefaultRetryPolicy() retry.Policy {
	return retry.NewRetryPolicy(retryAttempt)
}
