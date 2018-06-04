package business

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"strings"

	"github.com/syariatifaris/arkeus/core/fsm/webhook"
	"github.com/syariatifaris/arkeus/core/log/tokolog"
	"github.com/syariatifaris/arkeus/core/retry"
	"github.com/syariatifaris/arkeus/core/validation"
	"github.com/syariatifaris/arkeus/module/fsm/entity"
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

//GetFsmClientErrorMessage gets the client response operation cast
func (*BusinessModel) GetFsmClientErrorMessage(result interface{}) string {
	if result == nil {
		return ""
	}

	if wr, ok := result.(webhook.Response); ok {
		if fr, ok := wr.Result.(entity.FsmOperationResponse); ok {
			return strings.ToLower(fr.ErrMessage)
		}
	}

	return ""
}

//Log success info for business create / update order
func LogSuccessInfo(op, lp string, data interface{}) {
	var appID int64
	var appIDS []int64

	if fr, ok := data.(entity.FsmOperationResponse); ok {
		appID = fr.AppID
		appIDS = fr.AppIDs
	} else {
		if wr, ok := data.(webhook.Response); ok {
			if wr.Result != nil {
				if fr, ok := wr.Result.(entity.FsmOperationResponse); ok {
					appID = fr.AppID
					appIDS = fr.AppIDs
				}
			}
		}
	}

	msg := fmt.Sprintf("operation:%s success, app_id:[%+v]/app_id:[%+v]", op, appID, appIDS)
	tokolog.DEBUG.Println(lp, msg)
}
