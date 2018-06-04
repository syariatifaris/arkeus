package rubyistcb

import (
	"context"
	"errors"
	"time"

	"github.com/rubyist/circuitbreaker"
	"github.com/syariatifaris/arkeus/core/safeguard/common"
)

var (
	ErrorSafeguardTripped = errors.New("safeguard tripped, unsafe operation")
	ErrorSafeguardTimeout = errors.New("safeguard timeout")
)

const errBreakerOpenMsg = "breaker open"
const errBreakerTimeout = "breaker time out"

//breaker implementation

//CreateRateCb create rate circuit breaker
func CreateRateCb(errorRate float64, minSampling int64, timeout time.Duration, isEnabled bool) *RateCircuitBreakerImpl {
	cb := circuit.NewRateBreaker(errorRate, minSampling)
	return &RateCircuitBreakerImpl{
		cb:        cb,
		duration:  timeout,
		isEnabled: isEnabled,
	}
}

//RateCircuitBreakerImpl structure
type RateCircuitBreakerImpl struct {
	cb        *circuit.Breaker
	duration  time.Duration
	isEnabled bool
}

//Do runs the circuit breaker stub
func (rcb *RateCircuitBreakerImpl) Func(ctx context.Context, sf common.GuardFunc) error {
	var err, sfErr error
	if rcb.isEnabled {
		err = rcb.cb.Call(func() error {
			skip, err := sf(ctx)
			if skip {
				sfErr = err
				return nil
			}

			return err
		}, rcb.duration)
	} else {
		_, sfErr = sf(ctx)
	}

	if sfErr != nil {
		return sfErr
	}

	if err != nil {
		if err.Error() == errBreakerOpenMsg {
			return ErrorSafeguardTripped
		} else if err.Error() == errBreakerTimeout {
			return ErrorSafeguardTimeout
		}
	}

	return err
}

//Fail simply returns false flag for convention only
func (rcb *RateCircuitBreakerImpl) Fail() bool {
	return false
}

//Pass simply returns true flag for convention only
func (rcb *RateCircuitBreakerImpl) Pass() bool {
	return true
}
