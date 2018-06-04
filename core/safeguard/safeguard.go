package safeguard

import (
	"context"
	"time"

	"net/http"

	"github.com/syariatifaris/arkeus/core/safeguard/common"
	"github.com/syariatifaris/arkeus/core/safeguard/rubyistcb"
	"github.com/syariatifaris/arkeus/core/safeguard/tollbooth"
)

//NewRateSafeguard creates new safeguard with error rate context
func NewRateSafeguard(opt RateSafeguardConfig) Safeguard {
	return rubyistcb.CreateRateCb(opt.TripErrorRate, opt.MinSampling,
		time.Second*time.Duration(opt.TimeoutSecond), opt.IsEnabled)
}

//RateSafeguardConfig structure
type RateSafeguardConfig struct {
	MinSampling   int64
	TripErrorRate float64
	TimeoutSecond int64
	IsEnabled     bool
}

//Safeguard operation contract
type Safeguard interface {
	Func(ctx context.Context, safeguardFunc common.GuardFunc) error
	Fail() bool
	Pass() bool
}

//NewRequestLimiter creates new request limiter
func NewRequestLimiter(cfg LimiterConfig) RequestLimiter {
	return tollbooth.CreateHttpHandlerLimiter(cfg.MaxRequest, cfg.IsEnabled, cfg.RejectionHandler)
}

//RequestLimiter contract
type RequestLimiter interface {
	ProtectOverRequest(http.Handler) http.Handler
}

//LimiterConfig structure
type LimiterConfig struct {
	MaxRequest       float64
	IsEnabled        bool
	RejectionHandler http.HandlerFunc
}
