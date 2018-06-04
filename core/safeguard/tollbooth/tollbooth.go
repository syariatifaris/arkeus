package tollbooth

import (
	"net/http"

	"github.com/didip/tollbooth"
)

//CreateHttpHandlerLimiter creates http handler limiter instance
func CreateHttpHandlerLimiter(max float64, enabled bool,
	rejectFunc http.HandlerFunc) *TollboothImpl {
	return &TollboothImpl{
		maxRequest: max,
		isEnabled:  enabled,
		rejectFunc: rejectFunc,
	}
}

//TollboothImpl structure
type TollboothImpl struct {
	maxRequest float64
	isEnabled  bool
	rejectFunc http.HandlerFunc
}

//ProtectOverRequest protect exceeding http request
func (t *TollboothImpl) ProtectOverRequest(handler http.Handler) http.Handler {
	if !t.isEnabled {
		return handler
	}

	limiter := tollbooth.NewLimiter(t.maxRequest, nil)
	limiter.SetOnLimitReached(t.rejectFunc)

	return tollbooth.LimitHandler(limiter, handler)
}
