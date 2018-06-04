package retry

import (
	"errors"
	"fmt"
)

var (
	ErrOnRetryOnly = errors.New("retry(ies) routine error, secondary success")
)

//Create new retry policy instance
//arg attempt - total maximum attempt
func NewRetryPolicy(attempt int) Policy {
	return &policyImpl{
		maxAttempt: attempt,
	}
}

//Retry attempt policy interface
type Policy interface {
	Execute(ActionFunc) error
	SetMaxAttempt(int)

	ExecuteWithPlan(*Plan) error
	SetClosing(func())
	Finally()
}

type Plan struct {
	OnRetrying ActionRetryFunc
	OnFailure  SecondaryActionFunc
}

//Retry policy implementation
type policyImpl struct {
	c          func()
	maxAttempt int
}

//Retry action function
type ActionFunc func() (retry bool, err error)

//Retry action with secondary handle
type ActionRetryFunc func(p Policy) (retry bool, err error)

type SecondaryActionFunc func() error

//Set the maximum attempt value
func (r *policyImpl) SetMaxAttempt(val int) {
	r.maxAttempt = val
}

//Execute a stub with retry policy
//arg ActionFunc, function to be retried
func (r *policyImpl) Execute(fn ActionFunc) error {
	var err error
	var next bool
	attempt := 1
	for {
		next, err = fn()
		if !next || err == nil {
			break
		}
		attempt++
		if attempt > r.maxAttempt {
			return errors.New(fmt.Sprintf("Retrying %d times. Error cause: %s", r.maxAttempt, err.Error()))
		}
	}
	return err
}

func (r *policyImpl) ExecuteWithPlan(plan *Plan) error {
	if plan == nil {
		return errors.New("plan cannot be nil")
	}

	var err error
	var next bool
	attempt := 1
	for {
		next, err = plan.OnRetrying(r)
		if !next || err == nil {
			break
		}
		attempt++
		if attempt > r.maxAttempt {
			if plan.OnFailure != nil {
				retErr := err.Error()
				err = plan.OnFailure()
				if err != nil {
					msg := fmt.Sprintf("Retrying %d times. Error cause: %s.", r.maxAttempt, retErr)
					return errors.New(fmt.Sprintf("%s Secondary routine also failed, err: %s", msg, err.Error()))
				}

				return ErrOnRetryOnly
			}

			return errors.New(fmt.Sprintf("Retrying %d times. Error cause: %s", r.maxAttempt, err.Error()))
		}
	}
	return err
}

func (r *policyImpl) SetClosing(closingAction func()) {
	r.c = closingAction
}

func (r *policyImpl) Finally() {
	if r.c != nil {
		r.c()
	}

	r.c = nil
}
